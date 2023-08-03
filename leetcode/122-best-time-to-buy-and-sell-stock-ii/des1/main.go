package des1

func maxProfit(prices []int) int {
	return calculate(prices, 0)
}
func calculate(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}
	max := 0
	for start := s; start < len(prices); start++ {
		maxProfit := 0
		for i := start + 1; i < len(prices); i++ {
			if prices[start] < prices[i] {
				profit := calculate(prices, i+1) + prices[i] - prices[start]
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}
		if maxProfit > max {
			max = maxProfit
		}
	}
	return max
}
