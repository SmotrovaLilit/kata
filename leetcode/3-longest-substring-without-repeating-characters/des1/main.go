package des1

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start := 0
	set := make(map[byte]int) // contain the last index of the char
	for i := 0; i < len(s); i++ {
		if _, ok := set[s[i]]; ok {
			if set[s[i]] >= start {
				start = set[s[i]] + 1
			}
		}
		set[s[i]] = i
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}
