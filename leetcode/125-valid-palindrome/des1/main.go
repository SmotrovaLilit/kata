package des1

import (
	"regexp"
	"strings"
)

var (
	alphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func isPalindrome(s string) bool {
	s = alphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
