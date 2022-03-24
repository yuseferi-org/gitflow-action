package main

import (
	"context"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	//os.Setenv("GITHUB_TOKEN", "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")
	//orgName := "yuseferi-test-org"
	token := os.Getenv("GITHUB_TOKEN")
	orgName := os.Getenv("ORG_NAME")
	newRepoName := os.Getenv("NEW_REPO_NAME")
	githubUsername := os.Getenv("GITHUB_USERNAME")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// create a new private repository named "foo"
	repo := &github.Repository{
		Name:    github.String(newRepoName),
		Private: github.Bool(true),
	}
	_, _, err := client.Repositories.Create(ctx, orgName, repo)
	if err != nil {
		panic(err)
	}
	opt := &github.RepositoryAddCollaboratorOptions{Permission: "push"}
	_, _, err = client.Repositories.AddCollaborator(ctx, orgName, newRepoName, githubUsername, opt)
	if err != nil {
		panic(err)
	}

}
