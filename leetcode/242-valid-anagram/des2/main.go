package des1

import (
	"strconv"
	"strings"
)

func isAnagram(s string, t string) bool {
	return convertToKey(s) == convertToKey(t)
}

func convertToKey(s string) string {
	var count [26]int
	for _, v := range s {
		count[v-'a']++
	}
	buf := strings.Builder{}
	for _, v := range count {
		buf.WriteString("#" + strconv.Itoa(v))
	}
	return buf.String()
}
