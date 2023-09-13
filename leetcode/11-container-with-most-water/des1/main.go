package des1

func maxArea(height []int) int {
	r := 0
	min := 0
	for i, j := 0, len(height)-1; i < j; {
		min = height[i]
		if height[j] < min {
			min = height[j]
		}
		if r < min*(j-i) {
			r = min * (j - i)
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return r
}
