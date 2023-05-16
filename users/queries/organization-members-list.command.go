package queries

import (
	"context"
	"github.com/google/go-github/v52/github"
)

func GetOrganizationMembers(ctx context.Context, client *github.Client, org string) ([]*github.User, error) {
	opt := &github.ListMembersOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allUsers []*github.User
	for {
		users, resp, err := client.Organizations.ListMembers(ctx, org, opt)
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
