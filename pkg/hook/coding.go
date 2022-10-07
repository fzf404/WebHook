package hook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

// Coding Event Type
const (
	CodingPR   = "merge request"
	CodingPush = "push"
	CodingPing = "ping"
)

// Validate Coding Secret
func ValidateCoding(r *http.Request, secret string) bool {

	// Get Signature
	sign := r.Header.Get("X-Coding-Signature") // Read Signature
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

// Handle Coding Event
func HandleCoding(r *http.Request) (string, error) {

	// Get Event Type
	event := r.Header.Get("X-Coding-Event")

	switch event {
	case CodingPR:
		return PR, nil
	case CodingPush:
		return Push, nil
	case CodingPing:
		return Ping, nil
	}

	return event, errors.New("unknow event")
}
