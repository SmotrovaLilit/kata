package des1

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			next := dp[j] + 1
			if next > dp[i] {
				dp[i] = next
			}
		}
		if dp[i] > maxCount {
			maxCount = dp[i]
		}
	}
	return maxCount
}
