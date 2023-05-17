package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v52/github"
	"github.com/torinasakura/github/config"
	fileCommands "github.com/torinasakura/github/file/commands"
	issueCommands "github.com/torinasakura/github/issues/commands"
	userCommands "github.com/torinasakura/github/users/commands"
	userQueries "github.com/torinasakura/github/users/queries"
	"golang.org/x/oauth2"
	"log"
	"strings"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	teamMembers, err := userQueries.GetOrganizationTeamMembers(ctx, client, cfg.GithubOrg, cfg.GithubOrgTeamID)
	if err != nil {
		log.Fatalf("Could not get organization team members: %v", err)
	}

	for _, user := range teamMembers {
		repoName := fmt.Sprintf("%s.github.io", *user.Login)
		err = userCommands.CreateRepoForUser(ctx, client, user, repoName)
		if err != nil {
			log.Fatalf("Could not create repository: %v", err)
		}

		content, err := fileCommands.ReadMarkdownFileFromRepo(ctx, client.Repositories, cfg.GithubOrg, cfg.GithubTasksRepo, "README.md")
		if err != nil {
			log.Fatalf("Could not read file from repo: %v", err)
		}

		issueContent := strings.Split(content, "\n\n")
		for i, issue := range issueContent {
			err = issueCommands.CreateIssueFromMarkdown(ctx, client, cfg.GithubOrg, repoName, fmt.Sprintf("Task %d", i+1), issue)
			if err != nil {
				log.Fatalf("Could not create issue from markdown: %v", err)
			}
		}
	}
}
