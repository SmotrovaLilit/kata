package des1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			previousDp := dp[i-coins[j]]
			if previousDp == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] > previousDp+1 {
				dp[i] = previousDp + 1
			}

		}
	}
	return dp[amount]
}
