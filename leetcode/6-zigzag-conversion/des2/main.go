package des1

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	charsInSection := numRows + (numRows - 2)
	answer := strings.Builder{}
	for currentRaw := 0; currentRaw < numRows; currentRaw++ {
		index := currentRaw
		for index < len(s) {
			answer.WriteString(string(s[index]))
			if currentRaw != 0 && currentRaw != numRows-1 {
				charsInBetween := charsInSection - currentRaw*2
				secondIndex := index + charsInBetween
				if secondIndex < len(s) {
					answer.WriteString(string(s[secondIndex]))
				}
			}
			index += charsInSection
		}
	}
	return answer.String()
}
