package des1

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(strings.TrimSpace(s), " ")
	if len(pattern) != len(words) {
		return false
	}
	usedChar := make(map[byte]string)
	usedWord := make(map[string]struct{})
	for i := 0; i < len(pattern); i++ {
		if _, ok := usedChar[pattern[i]]; !ok {
			if _, ok := usedWord[strings.TrimSpace(words[i])]; ok {
				return false
			}
			usedChar[pattern[i]] = strings.TrimSpace(words[i])
			usedWord[strings.TrimSpace(words[i])] = struct{}{}
			continue
		}
		if usedChar[pattern[i]] != strings.TrimSpace(words[i]) {
			return false
		}
	}
	return true
}
