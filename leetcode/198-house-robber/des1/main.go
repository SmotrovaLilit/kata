package des1

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sumForNext := nums[len(nums)-1]
	sumForNextPlusOne := 0
	for i := len(nums) - 2; i >= 0; i-- {
		sumForNextPlusOne, sumForNext = sumForNext, max(sumForNext, sumForNextPlusOne+nums[i])
	}
	return sumForNext
}
