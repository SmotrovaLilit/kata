package des1

func isHappy(n int) bool {
	used := make(map[int]struct{})
	for n != 1 {
		used[n] = struct{}{}
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
		if _, ok := used[n]; ok {
			return false
		}
	}
	return true
}
