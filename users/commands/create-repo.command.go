package users

import (
	"context"
	"fmt"
	"github.com/google/go-github/v52/github"
)

func CreateRepoForUser(ctx context.Context, client *github.Client, user *github.User) error {
	repoName := *user.Login + "github.io"
	repo := &github.Repository{
		Name: github.String(repoName),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return fmt.Errorf("error creating repository for user %v: %w", *user.Login, err)
	}
	fmt.Println("Successfully created repository for user:", *user.Login)
	return nil
}
