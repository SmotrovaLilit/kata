package des1

import (
	"sort"
	"strings"
)

func topKFrequent(words []string, k int) []string {
	countWords := make(map[string]int)
	for _, v := range words {
		countWords[v]++
	}
	keys := make([]string, 0, len(countWords))
	for i, _ := range countWords {
		keys = append(keys, i)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if countWords[keys[i]] == countWords[keys[j]] {
			return strings.Compare(keys[i], keys[j]) < 0
		}

		return countWords[keys[i]] > countWords[keys[j]]
	})
	return keys[0:k]
}
