package des1

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, v := range strs {
		chars := strings.Split(v, "")
		sort.Strings(chars)
		key := strings.Join(chars, "")
		anagrams[key] = append(anagrams[key], v)
	}
	return convertMapToArray(anagrams)
}

func convertMapToArray(in map[string][]string) (out [][]string) {
	out = make([][]string, len(in), len(in))
	i := 0
	for _, v := range in {
		out[i] = v
		i++
	}
	return out
}
