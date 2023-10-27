package client

import (
	"context"
	"github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
	"os"
)

func NewClient() *github.Client {
	//repoPath := "/Users/kornik/repos/aerlingus/github/infrastructure-environments/"

	// Replace with the base and head commit hashes of the pull request.
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN environment variable not set")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
