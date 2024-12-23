package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/manifoldco/promptui"
	"github.com/mfridman/bestgo/gen/go/api"
	"github.com/mfridman/bestgo/gen/go/datapb"
	"github.com/mfridman/cli"
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
	root := &cli.Command{
		Name:        "bestgo",
		Usage:       "bestgo [flags]",
		Description: "bestgo is a CLI tool to get the star history of a Go repository",
		Flags: cli.FlagSetFunc(func(fs *flag.FlagSet) {
			fs.String("repo", "", "Full repository name. e.g., bufbuild/buf")
			fs.String("i", "year", "The grouping interval [day,month,quarter,year]")
		}),
		Exec: func(ctx context.Context, s *cli.State) error {
			var (
				repoName = cli.GetFlag[string](s, "repo")
				interval = cli.GetFlag[string](s, "i")

				client = api.NewClient(apiURL)
			)

			metrics, err := fetchRepoMetricsWithRetryPrompt(
				client,
				strings.TrimSpace(repoName),
				convertToTimeInterval(interval),
			)
			if err != nil {
				return fmt.Errorf("failed to get repo metrics from api: %w", err)
			}
			var points []point
			for _, m := range metrics.GetData() {
				points = append(points, point{date: m.X, count: m.Y})
			}

			fmt.Fprintln(s.Stdout, plot(points))
			fmt.Fprintf(s.Stdout, "Repository: %v has %d ⭐️ stars total\n",
				lipgloss.NewStyle().Foreground(lipgloss.Color("68")).Render(metrics.GetRepoName()),
				metrics.GetTotalStars(),
			)
			return nil
		},
	}

	if err := cli.ParseAndRun(context.Background(), root, os.Args[1:], nil); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
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
		name := repoName
		if name == "" {
			name = "unspecified"
		}
		fmt.Printf("Could not match repository: %s\n\n", name)
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
