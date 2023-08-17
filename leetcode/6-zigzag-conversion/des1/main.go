package des1

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, numRows)
	charsInSection := numRows + (numRows - 2)
	for i := 0; i < len(s); i++ {
		rowNumber := i % (charsInSection)
		if rowNumber < numRows {
			res[rowNumber] = append(res[rowNumber], s[i])
			continue
		}
		res[charsInSection-rowNumber] = append(res[charsInSection-rowNumber], s[i])
	}
	result := strings.Builder{}
	for _, v := range res {
		result.Write(v)
	}
	return result.String()
}
