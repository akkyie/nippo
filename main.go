package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akkyie/nippo/api"
	"github.com/akkyie/nippo/nippo"
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
	now := time.Now()

	client := api.NewClient(ctx, env.GithubAccessToken)

	username, err := client.GetUsername(ctx)
	if err != nil {
		log.Fatalf("failed to get current user: %v", err)
	}

	issues, err := client.ListIssues(ctx, username, now)
	if err != nil {
		log.Fatalf("failed to list issues: %v", err)
	}

	nippo := nippo.MakeNippo(issues, now)
	fmt.Print(nippo)
}
