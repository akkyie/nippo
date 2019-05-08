package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akkyie/nippo/api"
	"github.com/kelseyhightower/envconfig"
)

// Env is environment variables needed to run
type Env struct {
	GithubAccessToken string `required:"true" split_words:"true"`
}

func main() {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("failed to read environment variables: %v", err)
	}

	ctx := context.Background()

	client := api.NewClient(ctx, env.GithubAccessToken)

	username, err := client.GetUsername(ctx)
	if err != nil {
		log.Fatalf("failed to get current user: %v", err)
	}

	issues, err := client.ListIssues(ctx, username)
	if err != nil {
		log.Fatalf("failed to list issues: %v", err)
	}

	nippo := makeNippo(issues)
	fmt.Print(nippo)
}

func makeNippo(issues []api.Issue) string {
	today := time.Now().Format("2006-01-02")

	issueList := ""
	for _, issue := range issues {
		issueList += fmt.Sprintf("• %s %s\n", issue.Title, issue.URL)
	}

	template := `📅 日報 %s
*今日やること*
• …

*昨日やったこと*
• …
%s

*業務で気づいたこと*
• …
`

	return fmt.Sprintf(template, today, issueList)
}
