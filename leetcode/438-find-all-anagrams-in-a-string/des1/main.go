package des1

func findAnagrams(s string, p string) []int {
	pCharCount := convertToCharCounts(p)
	res := make([]int, 0)
	for i, v := range s {
		if pCharCount[v-'a'] == 0 {
			continue
		}
		var substring string
		if i+len(p) < len(s) {
			substring = s[i : i+len(p)]
		} else {
			substring = s[i:]
		}
		if isAnagram(pCharCount, substring) {
			res = append(res, i)
		}
	}
	return res
}

func isAnagram(pCharCount [26]int, s string) bool {
	var sCharCount [26]int
	for _, v := range s {
		if pCharCount[v-'a'] == 0 {
			return false
		}
		sCharCount[v-'a']++
	}
	for i, v := range sCharCount {
		if pCharCount[i] != v {
			return false
		}
	}
	return true
}

func convertToCharCounts(p string) [26]int {
	var res [26]int
	for _, v := range p {
		res[v-'a']++
	}
	return res
}
