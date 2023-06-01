package des1

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	var complement int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		// twoSum https://leetcode.com/problems/two-sum/solutions/  Two-pass Hash Table
		numsIndexes := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			complement = 0 - nums[i] - nums[j]
			_, ok := numsIndexes[complement]
			if ok {
				result = append(result, []int{
					nums[i], nums[j], complement,
				})
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}
			numsIndexes[nums[j]] = struct{}{}
		}

	}
	return result
}
