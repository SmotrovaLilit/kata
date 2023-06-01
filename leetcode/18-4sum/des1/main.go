package des1

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	for i, v1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if v1+nums[j]+nums[k]+nums[l] == target {
						set := make(map[int]struct{})
						set[v1] = struct{}{}
						set[nums[j]] = struct{}{}
						set[nums[k]] = struct{}{}
						set[nums[l]] = struct{}{}
						result = append(result, []int{v1, nums[j], nums[k], nums[l]})
					}
				}
			}
		}
	}
	return result
}
