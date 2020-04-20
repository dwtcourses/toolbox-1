package github

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func GetClient(ctx context.Context) *github.Client {
	var newClient *github.Client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	newClient = github.NewClient(tc)
	return newClient
}
