package des1

import "sort"

func threeSumSmaller(nums []int, target int) int {
	var count int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] >= target {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] >= target {
					break
				}
				count++
			}
		}
	}
	return count
}
