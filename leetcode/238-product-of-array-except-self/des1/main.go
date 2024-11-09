package des1

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = 1
		for j, v2 := range nums {
			if i != j {
				res[i] *= v2
			}
		}
	}
	return res
}
