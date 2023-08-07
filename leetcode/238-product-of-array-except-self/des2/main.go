package des1

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	product := 1
	for i, _ := range nums {
		product *= nums[i]
	}
	for i, _ := range nums {
		res[i] = product / nums[i]
	}
	return res
}
