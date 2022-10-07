package hook

import (
	"errors"
	"net/http"
)

const (
	GitlabPR      = "Merge Request Hook"
	GitlabTag     = "Tag Push Hook"
	GitlabPush    = "Push Hook"
	GitlabIssue   = "Issue Hook"
	GitlabComment = "Note Hook"
	GitlabRelease = "Release Hook"
)

// Validate Gitlab Secret
func ValidateGitlab(r *http.Request, secret string) bool {
	return secret == r.Header.Get("X-Gitlab-Token")
}

// Handle Gitlab Event
func HandleGitlab(r *http.Request) (string, error) {

	event := r.Header.Get("X-Gitlab-Event")

	switch event {
	case GitlabPR:
		return PR, nil
	case GitlabTag:
		return Tag, nil
	case GitlabPush:
		return Push, nil
	case GitlabIssue:
		return Issue, nil
	case GitlabComment:
		return Comment, nil
	case GitlabRelease:
		return Release, nil
	}

	return event, errors.New("unknow event")
}
