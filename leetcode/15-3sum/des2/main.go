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
			target := 0 - (nums[i] + nums[j])
			if target < nums[j] {
				continue
			}
			if target > nums[len(nums)-1] {
				continue
			}
			if k := search(nums[j+1:], target); k != -1 {
				result = append(result, []int{
					nums[i], nums[j], nums[k+j+1],
				})
			}
		}
	}
	return result
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
			continue
		}
		end = end - 1
	}
	return -1
}
