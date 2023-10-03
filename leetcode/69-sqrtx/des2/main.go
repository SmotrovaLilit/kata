package des1

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 2
	right := x / 2
	mid, num := 0, 0
	for left <= right {
		mid = left + (right-left)/2
		num = mid * mid
		if num == x {
			return mid
		}
		if num > x {
			right = mid - 1
			continue
		}
		if num < x {
			left = mid + 1
			continue
		}
	}
	return right
}
