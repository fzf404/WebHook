package hook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
)

const (
	GiteePR      = "Merge Request Hook"
	GiteeTag     = "Tag Push Hook"
	GiteePush    = "Push Hook"
	GiteeIssue   = "Issue Hook"
	GiteeComment = "Note Hook"
	GiteeRelease = "Release Hook"
)

// Validate Gitee Secret
func ValidateGitee(r *http.Request, secret string) bool {

	// Get Token
	token := r.Header.Get("X-Gitee-Token")

	if secret == token {
		return true
	}

	// Base64 Decode Secret
	sign := fmt.Sprintf("%s\n%s", r.Header.Get("X-Gitee-Timestamp"), secret)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(sign))
	mac_byte := mac.Sum(nil)
	mac_str := base64.StdEncoding.EncodeToString(mac_byte)

	return token == mac_str
}

// Handle Gitee Event
func HandleGitee(r *http.Request) (string, error) {

	event := r.Header.Get("X-Gitee-Event")

	switch event {
	case GiteePR:
		return PR, nil
	case GiteeTag:
		return Tag, nil
	case GiteePush:
		return Push, nil
	case GiteeIssue:
		return Issue, nil
	case GiteeComment:
		return Comment, nil
	case GiteeRelease:
		return Release, nil
	}

	return event, errors.New("unknow event")
}
