package config

import (
	"fmt"
	"os"
)

type Config struct {
	GithubToken     string
	GithubOrg       string
	GithubOrgTeamID string
	GithubTasksRepo string
}

func GetConfig() (*Config, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN not set in .env file")
	}

	githubOrg := os.Getenv("GITHUB_ORG")
	if githubOrg == "" {
		return nil, fmt.Errorf("GITHUB_ORG not set in .env file")
	}

	githubOrgTeamID := os.Getenv("GITHUB_ORG_TEAM_ID")
	if githubOrgTeamID == "" {
		return nil, fmt.Errorf("TEAM_ID not set in .env file")
	}

	githubTasksRepo := os.Getenv("GITHUB_TASKS_REPO")
	if githubTasksRepo == "" {
		return nil, fmt.Errorf("GITHUB_TASKS_REPO not set in .env file")
	}

	return &Config{
		GithubToken:     githubToken,
		GithubOrg:       githubOrg,
		GithubOrgTeamID: githubOrgTeamID,
		GithubTasksRepo: githubTasksRepo,
	}, nil
}
