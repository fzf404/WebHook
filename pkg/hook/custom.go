package hook

import (
	"errors"
	"net/http"
)

// Validate Custom Secret
func ValidateCustom(r *http.Request, secret string) bool {
	return secret == r.Header.Get("X-Custom-Token")
}

// Handle Custom Event
func HandleCustom(r *http.Request) (string, error) {

	event := r.Header.Get("X-Custom-Event")

	switch event {
	case PR:
		return PR, nil
	case Tag:
		return Tag, nil
	case Push:
		return Push, nil
	case Ping:
		return Ping, nil
	case Star:
		return Star, nil
	case Fork:
		return Fork, nil
	case Issue:
		return Issue, nil
	case Create:
		return Create, nil
	case Delete:
		return Delete, nil
	case Comment:
		return Comment, nil
	case Release:
		return Release, nil
	}

	return event, errors.New("unknow event")
}
