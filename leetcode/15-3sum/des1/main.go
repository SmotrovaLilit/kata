package des1

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != (i+1) && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k != (j+1) && nums[k] == nums[k-1] {
					continue
				}
				if 0 == nums[i]+nums[j]+nums[k] {
					result = append(result, []int{
						nums[i], nums[j], nums[k],
					})
				}

			}
		}
	}
	return result
}
