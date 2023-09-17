package des1

func getNext(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	for n != 1 && n != 4 {
		n = getNext(n)
	}
	return n == 1
}
