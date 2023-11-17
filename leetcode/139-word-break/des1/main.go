package des1

import (
	"slices"
	"strings"
)

type wordBreaking struct {
	cache    map[string]bool
	wordDict []string
	s        string
}

func newWordBreaking(wordDict []string, s string) *wordBreaking {
	return &wordBreaking{
		cache:    make(map[string]bool),
		wordDict: wordDict,
		s:        s,
	}

}

func (w *wordBreaking) wordBreak(s string) bool {
	if v, ok := w.cache[s]; ok {
		return v
	}
	for i := 0; i < len(w.wordDict); i++ {
		if !strings.Contains(s, w.wordDict[i]) {
			continue
		}
		stringWords := strings.Split(s, w.wordDict[i])
		isFind := true
		for j := 0; j < len(stringWords); j++ {
			if stringWords[j] == "" {
				continue
			}
			if !w.wordBreak(stringWords[j]) {
				isFind = false
			}
		}
		if isFind {
			w.cache[s] = true
			return true
		}
	}
	w.cache[s] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	slices.SortFunc(wordDict, func(i, j string) int {
		if len(i) > len(j) {
			return -1
		} else if len(i) < len(j) {
			return 1
		} else {
			return 0
		}
	})
	w := newWordBreaking(wordDict, s)
	return w.wordBreak(s)
}
