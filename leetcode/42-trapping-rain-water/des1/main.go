package des1

func trap(height []int) int {
	water := 0
	for i := 1; i < len(height)-1; i++ {
		left := 0
		right := 0
		// Search the left part for max bar size
		for j := i; j >= 0; j-- {
			if height[j] > left {
				left = height[j]
			}
		}
		// Search the right part for max bar size
		for j := i; j < len(height); j++ {
			if height[j] > right {
				right = height[j]
			}
		}
		if left > right {
			water += right - height[i]
		} else {
			water += left - height[i]
		}
	}
	return water
}
