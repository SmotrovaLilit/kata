package des1

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxRobbedAmount := make([]int, len(nums)+1)
	maxRobbedAmount[0] = 0
	maxRobbedAmount[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		maxRobbedAmount[i] = max(maxRobbedAmount[i+1], maxRobbedAmount[i+2]+nums[i])
	}
	return maxRobbedAmount[0]
}
