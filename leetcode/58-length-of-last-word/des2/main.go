package des1

func lengthOfLastWord(s string) int {
	result := 0
	i := len(s)
	for i > 0 {
		i--
		if s[i] != ' ' {
			result++
		} else if result > 0 {
			return result
		}

	}
	return result
}
