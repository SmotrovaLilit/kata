package des1

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	var start, end, target, sum int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		// twoSum2 https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/solutions/127822/two-sum-ii-input-array-is-sorted/
		start = i + 1
		end = len(nums) - 1
		target = 0 - nums[i]
		for start < end {
			sum = nums[start] + nums[end]
			if target == sum {
				result = append(result, []int{
					nums[i], nums[start], nums[end],
				})
				start++
				end--
				// exclude repeats
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				continue
			}
			if sum < target {
				start++
				continue
			}
			end--
		}
	}
	return result
}
