package commands

import (
	"context"
	"encoding/base64"
	"github.com/google/go-github/v52/github"
)

func ReadMarkdownFileFromRepo(ctx context.Context, client *github.Client, owner, repo, path string) (string, error) {
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(*fileContent.Content)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
