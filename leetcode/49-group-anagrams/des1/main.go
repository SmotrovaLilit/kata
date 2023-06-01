package des1

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	alreadyFoundAnagrams := make(map[int]struct{})
	for i := 0; i < len(strs); i++ {
		if _, ok := alreadyFoundAnagrams[i]; ok {
			continue
		}
		res = append(res, []string{strs[i]})

		for j := i + 1; j < len(strs); j++ {
			if _, ok := alreadyFoundAnagrams[j]; ok {
				continue
			}
			if isAnagram(strs[i], strs[j]) {
				alreadyFoundAnagrams[j] = struct{}{}
				res[len(res)-1] = append(res[len(res)-1], strs[j])
			}
		}

	}
	return res
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1MapCount := make(map[int32]int, len(s1))
	for _, v := range s1 {
		s1MapCount[v]++
	}
	s2MapCount := make(map[int32]int, len(s2))
	for _, v := range s2 {
		if _, ok := s1MapCount[v]; !ok {
			return false
		}
		s2MapCount[v]++
	}
	if len(s1MapCount) != len(s2MapCount) {
		return false
	}
	for k, v := range s1MapCount {
		if s2MapCount[k] != v {
			return false
		}
	}
	return true
}

func compareSets(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Map := make(map[int]struct{}, len(s1))
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]] = struct{}{}
	}
	for i := 0; i < len(s2); i++ {
		if _, ok := s1Map[s2[i]]; !ok {
			return false
		}
	}
	return true
}
