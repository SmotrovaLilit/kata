package des1

import "math"

func minSubArrayLen(target int, nums []int) int {
	start := 0
	sum := 0
	minSubArrayLen := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			minSubArrayLen = min(minSubArrayLen, i-start+1)
			sum -= nums[start]
			start++
		}
	}
	if minSubArrayLen == math.MaxInt {
		return 0
	}
	return minSubArrayLen
}
