package des1

import (
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, v := range strs {
		key := representKeyAsCountOfLetters(v)
		anagrams[key] = append(anagrams[key], v)
	}
	return convertMapToArray(anagrams)
}

func representKeyAsCountOfLetters(str string) string {
	var count [26]int
	for _, v := range str {
		count[v-'a']++
	}
	builder := strings.Builder{}
	for _, c := range count {
		builder.WriteString("#" + strconv.Itoa(c))
	}
	return builder.String()
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
