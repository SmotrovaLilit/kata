package des1

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	usedFromS := [256]byte{}
	usedFromT := [256]byte{}
	for i := 0; i < len(s); i++ {
		if usedFromS[s[i]] == 0 {
			if usedFromT[t[i]] != 0 {
				return false
			}
			usedFromT[t[i]] = s[i]
			usedFromS[s[i]] = t[i]
		} else if usedFromS[s[i]] != t[i] {
			return false
		}
	}
	return true
}
