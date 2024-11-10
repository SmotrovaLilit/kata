package des1

func getSumOfTarget(list []int, target int) (int, int) {
	left, right := 0, len(list)-1
	for left < right {
		currentSum := list[left] + list[right]
		if currentSum == target {
			return left, right
		}
		if currentSum < target {
			left = left + 1
			continue
		}
		right = right - 1
	}
	return 0, 0
}
