package des1

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] >= target {
				break
			}
			count++
		}
	}
	return count
}
