package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/peterbourgon/ff/v3"
	"go.buf.build/demolab/template-twirp/mf192/bestofgo/api"
)

func main() {
	fs := flag.NewFlagSet("bestgo", flag.ExitOnError)
	var (
		repoName = fs.String("repo", "", "full repository name. Example: go-chi/chi (mandatory)")
		interval = fs.String("i", "year", "grouping interval. Supported: year, quarter, month")
	)
	if err := ff.Parse(fs, os.Args[1:]); err != nil {
		log.Fatalf("failed to parse flags: %v", err)
	}
	if *repoName == "" {
		fmt.Printf("error: repo cannot be empty. Example -repo go-chi/chi\n")
		fs.Usage()
		os.Exit(1)
	}
	var timeInterval api.TimeIntervalType
	switch *interval {
	case "month":
		timeInterval = api.TimeIntervalType_TIME_INTERVAL_TYPE_MONTH
	case "quarter":
		timeInterval = api.TimeIntervalType_TIME_INTERVAL_TYPE_QUARTER
	default:
		timeInterval = api.TimeIntervalType_TIME_INTERVAL_TYPE_YEAR
	}

	client := api.NewAPIServiceProtobufClient("https://api.bestofgo.dev", http.DefaultClient)
	resp, err := client.GetRepoMetrics(context.Background(), &api.GetRepoMetricsRequest{
		RepoName:     *repoName,
		TimeInterval: timeInterval,
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("error: failed to get repo metrics from api: %w", err))
	}
	var points []point
	for _, m := range resp.GetData() {
		points = append(points, point{date: m.X, count: m.Y})
	}
	fmt.Println(plot(points))
	fmt.Printf("Repository: %v has %d ⭐️ stars total\n", *repoName, resp.GetTotalStars())
}

type point struct {
	date  string
	count int32
}

func plot(points []point) string {
	var max int32
	for _, p := range points {
		if p := p.count; p > max {
			max = p
		}
	}
	var b strings.Builder
	for i := 0; i < len(points); i++ {
		var l int32
		if max > 0 {
			l = (points[i].count*50 + max/2) / max
		}
		b.WriteString(fmt.Sprintf("%s [%v]\t|%v\n",
			points[i].date,
			points[i].count,
			strings.Repeat(`■`, int(l))))
	}
	return b.String()
}
