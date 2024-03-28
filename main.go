package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/manifoldco/promptui"
	"github.com/mfridman/bestgo/gen/go/api"
	"github.com/mfridman/bestgo/gen/go/datapb"
	"github.com/peterbourgon/ff/v3"
	"github.com/ryanuber/columnize"
)

const (
	apiURL = "https://api.bestofgo.dev"
)

var (
	re = regexp.MustCompile(`(?m)\[(.*?)\s*\]`)
)

func convertToTimeInterval(interval string) datapb.TimeIntervalType {
	switch interval {
	case "day":
		return datapb.TimeIntervalType_TIME_INTERVAL_TYPE_DAY
	case "month":
		return datapb.TimeIntervalType_TIME_INTERVAL_TYPE_MONTH
	case "quarter":
		return datapb.TimeIntervalType_TIME_INTERVAL_TYPE_QUARTER
	default:
		return datapb.TimeIntervalType_TIME_INTERVAL_TYPE_YEAR
	}
}

func main() {
	fs := flag.NewFlagSet("bestgo", flag.ExitOnError)
	var (
		repoName = fs.String("repo", "", "full repository name. Example: bufbuild/buf (mandatory)")
		interval = fs.String("i", "year", "grouping interval. Supported: year, quarter, month, day")
	)
	if err := ff.Parse(fs, os.Args[1:]); err != nil {
		log.Fatalf("failed to parse flags: %v", err)
	}
	if *repoName == "" {
		fmt.Printf("error: repo cannot be empty. Example -repo bufbuild/buf\n")
		fs.Usage()
		os.Exit(1)
	}

	client := api.NewClient(apiURL)

	metrics, err := fetchRepoMetricsWithRetryPrompt(
		client,
		strings.TrimSpace(*repoName),
		convertToTimeInterval(*interval),
	)
	if err != nil {
		log.Fatalln(fmt.Errorf("error: failed to get repo metrics from api: %w", err))
	}
	var points []point
	for _, m := range metrics.GetData() {
		points = append(points, point{date: m.X, count: m.Y})
	}
	fmt.Println(plot(points))
	fmt.Printf("Repository: %v has %d ⭐️ stars total\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("68")).Render(metrics.GetRepoName()),
		metrics.GetTotalStars(),
	)
}

type point struct {
	date  string
	count int32
}

// could also be done directly in sql.
// https://tapoueh.org/blog/2014/02/postgresql-aggregates-and-histograms/?utm_source=pocket_mylist
func plot(points []point) string {
	var max int32
	for _, p := range points {
		if p.count > max {
			max = p.count
		}
	}
	var b strings.Builder
	for i := 0; i < len(points); i++ {
		var l int32
		if max > 0 {
			l = (points[i].count*30 + max/2) / max
		}
		b.WriteString(fmt.Sprintf("%s [%v]%s|%v\n",
			points[i].date,
			points[i].count,
			strings.Repeat(" ", addPadding(points[i].count, max)),
			strings.Repeat(`■`, int(l))))
	}
	return b.String()
}

func addPadding(count, max int32) int {
	m := len(strconv.Itoa(int(max)))
	c := len(strconv.Itoa(int(count)))
	if c >= m {
		return 0
	}
	return m - c
}

func fetchRepoMetricsWithRetryPrompt(
	client *api.Client,
	repoName string,
	timeInterval datapb.TimeIntervalType,
) (*api.GetRepoMetricsResponse, error) {
	resp, err := client.APIService.GetRepoMetrics(context.Background(), &api.GetRepoMetricsRequest{
		RepoName:     repoName,
		TimeInterval: timeInterval,
	})
	if err != nil {
		return nil, err
	}
	if len(resp.GetSuggestions()) > 0 {
		fmt.Printf("Could not match repository: %s\n\n", repoName)
		prompt := promptui.Select{
			Label: "Was that a typo? Did you mean one of these",
		}
		var items []string
		for i, repo := range resp.GetSuggestions() {
			i++
			items = append(items, fmt.Sprintf("%d.| [|%s|] with %d|stars",
				i,
				repo.GetRepoFullName(),
				repo.GetRepoStargazers()),
			)
		}
		prompt.Items = strings.Split(columnize.SimpleFormat(items), "\n")

		_, result, err := prompt.Run()
		if err != nil {
			return nil, fmt.Errorf("failed on ui prompt: %w", err)
		}
		out := re.FindStringSubmatch(result)
		if len(out) == 2 {
			repoName = strings.TrimSpace(out[1])
			resp, err := client.APIService.GetRepoMetrics(context.Background(), &api.GetRepoMetricsRequest{
				RepoName:     repoName,
				TimeInterval: timeInterval,
			})
			if err != nil {
				return nil, err
			}
			return resp, nil
		}
		return nil, fmt.Errorf("failed to match repository: %v", repoName)
	}

	return resp, nil
}
