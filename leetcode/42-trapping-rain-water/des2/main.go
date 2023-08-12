package des1

func trap(height []int) int {
	water := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = leftMax[i-1]
		if height[i] > leftMax[i] {
			leftMax[i] = height[i]
		}
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = rightMax[i+1]
		if height[i] > rightMax[i] {
			rightMax[i] = height[i]
		}
	}
	for i := 1; i < len(height)-1; i++ {
		if leftMax[i] > rightMax[i] {
			water += rightMax[i] - height[i]
		} else {
			water += leftMax[i] - height[i]
		}
	}
	return water
}
