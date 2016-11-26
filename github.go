package main

import (
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func getGitHubClient() *github.Client {
	oauth2Token := &oauth2.Token{
		AccessToken: os.Getenv("GITHUB_API_TOKEN"),
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(oauth2Token))
	client := github.NewClient(oauthClient)

	return client
}
