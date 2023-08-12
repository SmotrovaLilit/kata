package des1

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	water := 0
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
			continue
		}
		if height[right] >= rightMax {
			rightMax = height[right]
		} else {
			water += rightMax - height[right]
		}
		right--
	}
	return water
}
