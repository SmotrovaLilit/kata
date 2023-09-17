package des1

func getNext(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

var cycleMembers = map[int]struct{}{
	4: {}, 16: {}, 37: {}, 58: {}, 89: {}, 145: {}, 42: {}, 20: {},
}

func isHappy(n int) bool {
	for n != 1 {
		if _, ok := cycleMembers[n]; ok {
			return false
		}
		n = getNext(n)
	}
	return n == 1
}
