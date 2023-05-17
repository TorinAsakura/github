package commands

import (
	"context"
	"encoding/base64"
	"github.com/google/go-github/v52/github"
)

type GitHubClient interface {
	GetContents(ctx context.Context, owner, repo, path string, opt *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error)
}

func ReadMarkdownFileFromRepo(ctx context.Context, client GitHubClient, owner, repo, path string) (string, error) {
	fileContent, _, _, err := client.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(*fileContent.Content)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
