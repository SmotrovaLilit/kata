package des1

import (
	"strings"
)

type wordBreaking struct {
	cache    map[int]bool
	wordDict []string
	s        string
}

func newWordBreaking(wordDict []string, s string) *wordBreaking {
	return &wordBreaking{
		cache:    make(map[int]bool),
		wordDict: wordDict,
		s:        s,
	}
}

func (w *wordBreaking) dp(i int) bool {
	if v, ok := w.cache[i]; ok {
		return v
	}
	s := w.s[:i+1]
	for j := 0; j < len(w.wordDict); j++ {
		if !strings.HasSuffix(s, w.wordDict[j]) {
			continue
		}
		newI := i - len(w.wordDict[j])
		if newI < 0 { // w.wordDict[j] == s
			w.cache[i] = true
			return true
		}
		if w.dp(newI) {
			w.cache[i] = true
			return true
		}
	}
	w.cache[i] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	w := newWordBreaking(wordDict, s)
	return w.dp(len(s) - 1)
}
