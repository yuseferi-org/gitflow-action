package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	token := os.Getenv("ADMIN_GITHUB_TOKEN")
	orgName := os.Getenv("ORG_NAME")
	newRepoName := os.Getenv("NEW_REPO_NAME")
	githubUsername := os.Getenv("GITHUB_USERNAME")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// create a new private repository
	repo := &github.Repository{
		Name:    github.String(newRepoName),
		Private: github.Bool(true),
	}
	repository, _, err := client.Repositories.Create(ctx, orgName, repo)
	if err != nil {
		panic(err)
	}
	opt := &github.RepositoryAddCollaboratorOptions{Permission: "push"}
	_, _, err = client.Repositories.AddCollaborator(ctx, orgName, newRepoName, githubUsername, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(repository.GetHTMLURL())
	fmt.Println("::set-output name=repo_url::test")

}
