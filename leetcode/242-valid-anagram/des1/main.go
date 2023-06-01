package des1

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	return convertToKey(s) == convertToKey(t)
}

func convertToKey(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
