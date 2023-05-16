package queries

import (
	"context"
	"github.com/google/go-github/v52/github"
)

func GetOrganizationTeamMembers(ctx context.Context, client *github.Client, org string, teamSlug string) ([]*github.User, error) {
	opt := &github.TeamListTeamMembersOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allUsers []*github.User
	for {
		users, resp, err := client.Teams.ListTeamMembersBySlug(ctx, org, teamSlug, opt)
		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, users...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allUsers, nil
}
