package des1

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		rotateOne(nums)
	}
}

func rotateOne(nums []int) {
	if len(nums) == 0 {
		return
	}
	last := nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = last
}
