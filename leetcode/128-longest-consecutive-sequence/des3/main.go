package des1

import "sort"

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	longest := 0
	for i := 0; i < len(nums); {
		start := nums[i]
		end := nums[i]
		for i < len(nums)-1 && (nums[i+1] == end+1 || nums[i+1] == end) {
			if nums[i+1] == end {
				i++
				continue
			}
			end++
			i++
		}
		if end-start+1 > longest {
			longest = end - start + 1
		}
		i++
	}

	return longest
}
