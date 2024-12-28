package github

import (
	"fmt"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/shufo/gh-issue-stats/internal/utils"
	"github.com/shufo/gh-issue-stats/pkg/types"
)

var debug bool

func SetDebug(d bool) {
	debug = d
}

func GetRepoInfo() (string, error) {
	stdOut, stdErr, err := gh.Exec("repo", "view", "--json", "nameWithOwner", "-q", ".nameWithOwner")
	if err != nil {
		return "", fmt.Errorf("%v", stdErr.String())
	}
	return stdOut.String()[:stdOut.Len()-1], nil
}

func FetchIssues() ([]types.Issue, error) {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create GitHub client: %v", err)
	}

	repo, err := GetRepoInfo()
	if err != nil {
		return nil, err
	}

	// First, get the total count of issues to calculate pages
	var totalCount int
	// path := fmt.Sprintf("repos/%s/issues?state=all&per_page=1", repo)
	path := fmt.Sprintf("search/issues?q=repo:%s", repo)
	response := struct {
		TotalCount int `json:"total_count"`
	}{}

	err = client.Get(path, &response)
	if err == nil {
		totalCount = response.TotalCount
	}

	// Create and start spinner
	perPage := 100
	totalPages := (totalCount + perPage - 1) / perPage

	if totalPages > 0 {
		utils.DebugPrintf("Total issues (including PRs): %d", totalCount)
		utils.DebugPrintf("Total pages: %d\n", totalPages)
	}

	utils.StartSpinner(" Fetching issues...")

	var allIssues []types.Issue

	utils.DebugPrintf("starting to fetch issues")

	for page := 1; page <= totalPages; page++ {
		if debug {
			utils.DebugPrintf("fetching issues (%d/%d)", page, totalPages)
		} else {
			utils.UpdateSpinnerSuffix(fmt.Sprintf(" Fetching issues... (%d/%d)", page, totalPages))
		}

		var pageIssues []types.Issue
		path := fmt.Sprintf("repos/%s/issues?state=all&per_page=%d&page=%d", repo, perPage, page)
		err := client.Get(path, &pageIssues)
		if err != nil {
			utils.StopSpinner()
			return nil, fmt.Errorf("failed to fetch issues: %v", err)
		}

		if len(pageIssues) == 0 {
			break
		}

		// Filter out pull requests and count issues
		issuesCount := 0
		for _, issue := range pageIssues {
			if issue.PullRequest == nil {
				allIssues = append(allIssues, issue)
				issuesCount++
			}
		}
		utils.DebugPrintf("fetched %d: found %d issues (total so far: %d)",
			page, issuesCount, len(allIssues))
	}

	// Stop spinner and clear the line
	if !debug {
		utils.StopSpinner()
	}

	utils.DebugPrintf("finished fetching issues (total: %d)", len(allIssues))
	return allIssues, nil
}
