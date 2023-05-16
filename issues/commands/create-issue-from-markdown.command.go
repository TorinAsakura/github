package commands

import (
	"context"
	"github.com/google/go-github/v52/github"
)

func CreateIssueFromMarkdown(ctx context.Context, client *github.Client, org, repo, title, body string) error {
	issue := &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	}
	_, _, err := client.Issues.Create(ctx, org, repo, issue)
	return err
}
