package des1

func romanToInt(s string) int {
	romanToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for k := 0; k < len(s); k++ {
		v := s[k]
		ch := string(v)
		if len(s) > k+1 {
			chNext := string(s[k+1])
			if ch == "I" && (chNext == "V" || chNext == "X") {
				res += romanToInt[chNext] - romanToInt[ch]
				k++
				continue
			}
			if ch == "X" && (chNext == "L" || chNext == "C") {
				res += romanToInt[chNext] - romanToInt[ch]
				k++
				continue
			}
			if ch == "C" && (chNext == "D" || chNext == "M") {
				res += romanToInt[chNext] - romanToInt[ch]
				k++
				continue
			}
		}
		res += romanToInt[ch]
	}
	return res
}
