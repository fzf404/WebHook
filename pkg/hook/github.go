package hook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	GithubPR      = "pull_request"
	GithubPush    = "push"
	GithubPing    = "ping"
	GithubStar    = "star"
	GithubFork    = "fork"
	GithubCreate  = "create"
	GithubDelete  = "delete"
	GithubIssue   = "issues"
	GithubComment = "issue_comment"
	GithubRelease = "release"
)

// Validate Github Secret
func ValidateGithub(r *http.Request, secret string) bool {

	// Get Signature
	sign := r.Header.Get("X-Hub-Signature") // Read Signature
	if len(sign) != 45 {
		return false
	}

	// HMAC Secret
	body, _ := ioutil.ReadAll(r.Body)         // Read Body
	mac := hmac.New(sha1.New, []byte(secret)) // New HMAC
	mac.Write(body)                           // Join Body
	mac_byte := mac.Sum(nil)                  // Sum HMAC
	mac_str := hex.EncodeToString(mac_byte)   // Encode HMAC

	return sign[5:] == mac_str
}

// Handle Github Event
func HandleGithub(r *http.Request) (string, error) {

	event := r.Header.Get("X-Github-Event")

	switch event {
	case GithubPR:
		return PR, nil
	case GithubPush:
		return Push, nil
	case GithubPing:
		return Ping, nil
	case GithubStar:
		return Star, nil
	case GithubFork:
		return Fork, nil
	case GithubIssue:
		return Issue, nil
	case GithubCreate:
		return Create, nil
	case GithubDelete:
		return Delete, nil
	case GithubComment:
		return Comment, nil
	case GithubRelease:
		return Release, nil
	}

	return event, errors.New("unknow event")
}
