package des1

import "math"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
