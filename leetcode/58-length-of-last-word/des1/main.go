package des1

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	// trim right
	for i >= 0 && s[i] == ' ' {
		i--
	}
	result := 0
	// count last word
	for i >= 0 && s[i] != ' ' {
		result++
		i--
	}
	return result
}
