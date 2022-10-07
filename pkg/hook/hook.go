package hook

import (
	"errors"
	"net/http"
	"strings"
)

// Hook Event
const (
	PR      = "pr"
	Tag     = "tag"
	Push    = "push"
	Ping    = "ping"
	Star    = "star"
	Fork    = "fork"
	Issue   = "issue"
	Create  = "create"
	Delete  = "delete"
	Comment = "comment"
	Release = "release"

	Unknow = "unknow"
)

// Platform
const (
	Github = "github"
	Gitee  = "gitee"
	Coding = "coding"
	Gitlab = "gitlab"
	Gitea  = "gitea"
	Custom = "custom"
)

// Platform UesrAgent
const (
	GithubUA = "GitHub"
	GiteeUA  = "git-oschina"
	CodingUA = "Coding"
	GitlabUA = "GitLab"
	GiteaUA  = "Go-http"
	CustomUA = "Custom"
)

// Judge Request Platform
func JudgePlatform(r *http.Request) (string, error) {

	// Get UserAgent
	ua := r.Header.Get("User-Agent")

	// Switch Request Platform
	switch {
	case strings.HasPrefix(ua, GithubUA):
		return Github, nil
	case strings.HasPrefix(ua, GiteeUA):
		return Gitee, nil
	case strings.HasPrefix(ua, CodingUA):
		return Coding, nil
	case strings.HasPrefix(ua, GitlabUA):
		return Gitlab, nil
	case strings.HasPrefix(ua, GiteaUA):
		return Gitea, nil
	}

	return Custom, nil
}

// Validate Platform Secret
func ValidatePlatform(r *http.Request, platform string, secret string) bool {

	switch platform {
	case Github:
		return ValidateGithub(r, secret)
	case Gitee:
		return ValidateGitee(r, secret)
	case Coding:
		return ValidateCoding(r, secret)
	case Gitlab:
		return ValidateGitlab(r, secret)
	case Gitea:
		return ValidateGitea(r, secret)
	case Custom:
		return ValidateCustom(r, secret)
	}

	return false
}

// Handle Platform Event
func HandlePlatform(r *http.Request, platform string) (string, error) {

	switch platform {
	case Github:
		return HandleGithub(r)
	case Gitee:
		return HandleGitee(r)
	case Coding:
		return HandleCoding(r)
	case Gitlab:
		return HandleGitlab(r)
	case Gitea:
		return HandleGitea(r)
	case Custom:
		return HandleCustom(r)
	}

	return Unknow, errors.New("unknow platform")
}
