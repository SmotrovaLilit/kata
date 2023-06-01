package des1

func findAnagrams(s string, p string) []int {
	pCharsCount := make(map[int32]int)
	sCharsCount := make(map[int32]int)
	res := make([]int, 0)
	for _, v := range p {
		pCharsCount[v]++
	}
	for i, sChar := range s {
		sCharsCount[sChar]++
		if i >= len(p) {
			k := int32(s[i-len(p)])
			sCharsCount[k]--
			if sCharsCount[k] == 0 {
				delete(sCharsCount, k)
			}
		}
		if isEqualMaps(pCharsCount, sCharsCount) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func isEqualMaps(m1 map[int32]int, m2 map[int32]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
