package commands

import (
	"context"
	"encoding/base64"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v52/github"
	"testing"
)

func TestReadMarkdownFileFromRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockGitHubClient(ctrl)
	fileContent := "test content"
	encodedContent := base64.StdEncoding.EncodeToString([]byte(fileContent))
	githubFileContent := &github.RepositoryContent{Content: &encodedContent}
	mockClient.EXPECT().GetContents(gomock.Any(), "owner", "repo", "path", gomock.Nil()).Return(githubFileContent, nil, nil, nil)

	ctx := context.Background()
	content, err := ReadMarkdownFileFromRepo(ctx, mockClient, "owner", "repo", "path")
	if err != nil {
		t.Errorf("ReadMarkdownFileFromRepo returned error: %v", err)
	}

	if content != fileContent {
		t.Errorf("Expected content '%s', but got '%s'", fileContent, content)
	}
}
