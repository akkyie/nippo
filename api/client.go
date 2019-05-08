package api

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client wraps GitHub API access
type Client struct {
	gh *github.Client
}

// Issue represents an issue or a pull request on GitHub
type Issue struct {
	Title string
	URL   string
}

// NewClient returns a GitHub API client authenticated with given access token
func NewClient(ctx context.Context, accessToken string) *Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	oauth2Client := oauth2.NewClient(ctx, tokenSource)
	return &Client{github.NewClient(oauth2Client)}
}

// GetUsername returns the authenticated user's login name
func (client *Client) GetUsername(ctx context.Context) (string, error) {
	result, _, err := client.gh.Users.Get(ctx, "")
	if err != nil {
		return "", err
	}
	return result.GetLogin(), nil
}

// ListIssues lists issues including pull requests updated since yesterday
func (client *Client) ListIssues(ctx context.Context, username string) ([]Issue, error) {
	options := &github.SearchOptions{Sort: "updated", Order: "asc"}
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	query := fmt.Sprintf("updated:>=%s involves:%s", yesterday, username)
	result, _, err := client.gh.Search.Issues(ctx, query, options)
	if err != nil {
		return nil, err
	}

	issues := make([]Issue, len(result.Issues))
	for i, issue := range result.Issues {
		issues[i] = Issue{Title: issue.GetTitle(), URL: issue.GetHTMLURL()}
	}

	return issues, nil
}
