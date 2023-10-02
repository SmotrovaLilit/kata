package des1

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 { // exclude 10, 20, 100, 1000, ...
		return false
	}
	revertedHalf := 0
	firstHalf := x
	for firstHalf > revertedHalf {
		revertedHalf = revertedHalf*10 + firstHalf%10
		firstHalf /= 10
	}

	return firstHalf == revertedHalf || firstHalf == revertedHalf/10
}
