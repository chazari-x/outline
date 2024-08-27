package outline

import (
	"strings"
)

// Outline represents an outline
type Outline struct {
	apiUrl     string
	certSha256 string
}

// NewOutline creates a new outline
//
// Example: NewOutline("https://xxx.xxx.xxx.xxx:xxxxx/xxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxx")
func NewOutline(apiUrl, certSha256 string) *Outline {
	return &Outline{apiUrl, strings.ToLower(certSha256)}
}
