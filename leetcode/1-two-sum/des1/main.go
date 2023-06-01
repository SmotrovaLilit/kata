package des1

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v1+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
