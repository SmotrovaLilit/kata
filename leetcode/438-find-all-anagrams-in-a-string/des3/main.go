package des1

func findAnagrams(s string, p string) []int {
	pCharsCount := convertToCharCounts(p)
	res := make([]int, 0)
	var sCharsCount [26]int
	for i, sChar := range s {
		sCharsCount[sChar-'a']++
		if i >= len(p) {
			sCharsCount[int32(s[i-len(p)])-'a']--
		}
		if isEqual(pCharsCount, sCharsCount) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func isEqual(m1 [26]int, m2 [26]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
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
