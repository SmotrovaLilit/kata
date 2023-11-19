package des1

func lengthOfLIS(nums []int) int {
	s := newSolution(nums)
	maxCount := 0
	for i := len(nums) - 1; i >= 0; i-- {
		next := s.dp(i)
		if next > maxCount {
			maxCount = next
		}
	}
	return maxCount
}

type solution struct {
	cache map[int]int
	nums  []int
}

func newSolution(nums []int) *solution {
	return &solution{
		cache: make(map[int]int, len(nums)),
		nums:  nums,
	}
}

func (s *solution) dp(i int) int {
	if v, ok := s.cache[i]; ok {
		return v
	}

	maxCount := 1
	for j := 0; j < i; j++ {
		if s.nums[j] >= s.nums[i] {
			continue
		}
		next := s.dp(j) + 1
		if next > maxCount {
			maxCount = next
		}
	}
	s.cache[i] = maxCount
	return maxCount
}
