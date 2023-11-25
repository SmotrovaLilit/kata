package des1

import "sort"

func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	longest := 0
	var num int
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		num = nums[i]
		start := num
		end := num
		for m[end+1] {
			end++
		}
		if end-start+1 > longest {
			longest = end - start + 1
		}
		i += end - start + 1
	}

	return longest
}
