package des1

func coinChange(coins []int, amount int) int {
	return newSolution(coins).dp(amount)
}

type solution struct {
	cache map[int]int
	coins []int
}

func newSolution(coins []int) *solution {
	return &solution{
		cache: make(map[int]int),
		coins: coins,
	}
}

func (s *solution) dp(amount int) int {
	if amount == 0 {
		return 0
	}
	if v, ok := s.cache[amount]; ok {
		return v
	}
	minimalCount := -1
	nextMin := -1
	for i := 0; i < len(s.coins); i++ {
		if amount < s.coins[i] {
			continue
		}
		if amount == s.coins[i] {
			minimalCount = 1
			break
		}
		nextMin = s.dp(amount - s.coins[i])
		if nextMin == -1 {
			continue
		}
		nextMin++
		if nextMin <= minimalCount || minimalCount == -1 {
			minimalCount = nextMin
		}
	}
	s.cache[amount] = minimalCount
	return minimalCount
}
