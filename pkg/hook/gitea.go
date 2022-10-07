package hook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	GiteaPR      = "pull_request"
	GiteaPush    = "push"
	GiteaCreate  = "create"
	GiteaDelete  = "delete"
	GiteaIssue   = "issues"
	GiteaComment = "issue_comment"
	GiteaRelease = "release"
)

// Validate Gitea Secret
func ValidateGitea(r *http.Request, secret string) bool {

	// Get Signature
	sign := r.Header.Get("X-Gitea-Signature") // Read Signature

	// HMAC Secret
	body, _ := ioutil.ReadAll(r.Body)           // Read Body
	mac := hmac.New(sha256.New, []byte(secret)) // New HMAC
	mac.Write(body)                             // Join Body
	mac_byte := mac.Sum(nil)                    // Sum HMAC
	mac_str := hex.EncodeToString(mac_byte)     // Encode HMAC

	return sign == mac_str
}

// Handle Gitea Event
func HandleGitea(r *http.Request) (string, error) {

	event := r.Header.Get("X-Gitea-Event")

	switch event {
	case GiteaPR:
		return PR, nil
	case GiteaPush:
		return Push, nil
	case GiteaCreate:
		return Create, nil
	case GiteaDelete:
		return Delete, nil
	case GiteaIssue:
		return Issue, nil
	case GiteaComment:
		return Comment, nil
	case GiteaRelease:
		return Release, nil
	}

	return event, errors.New("unknow event")
}
