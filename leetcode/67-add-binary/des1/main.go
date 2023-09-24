package des1

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""
	var carry, sum int
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i-- {
		sum = carry
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		carry = sum / 2
		result = fmt.Sprintf("%d%s", sum%2, result)
		j--
	}
	return result
}
