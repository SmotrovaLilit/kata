package des1

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}

	return maxProfit
}
