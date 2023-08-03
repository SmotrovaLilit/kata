package des1

func rotate(nums []int, k int) {
	count := 0
	for i := 0; count < len(nums); i++ {
		current := i
		prevValue := nums[i]
		for {
			next := (current + k) % len(nums)
			tempValue := nums[next]
			nums[next] = prevValue
			prevValue = tempValue
			current = next
			count++
			if i == current {
				break
			}
		}
	}
}
