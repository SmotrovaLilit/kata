package des1

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	res := strings.Builder{}
	v := ""
	for i := len(words) - 1; i >= 0; i-- {
		v = strings.TrimSpace(words[i])
		if v == "" {
			continue
		}
		res.WriteString(v)
		res.WriteString(" ")
	}
	return strings.TrimRight(res.String(), " ")
}
