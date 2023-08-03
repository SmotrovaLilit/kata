package des1

func maxProfit(prices []int) int {
	peak := prices[0]
	valley := prices[0]
	maxProfit := 0
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxProfit += peak - valley
	}
	return maxProfit
}
