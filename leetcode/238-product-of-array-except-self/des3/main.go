package des1

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	r := make([]int, len(nums))

	l[0] = 1
	for i := 1; i < len(l); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i, _ := range nums {
		nums[i] = l[i] * r[i]
	}
	return nums
}
