package des1

import "slices"

func lengthOfLIS(nums []int) int {
	sub := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			j, _ := slices.BinarySearch(sub, nums[i])
			sub[j] = nums[i]
		}
	}
	return len(sub)
}
