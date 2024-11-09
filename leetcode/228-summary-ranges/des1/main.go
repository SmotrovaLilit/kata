package des1

import "fmt"

func summaryRanges(nums []int) []string {
	var result []string
	var start int
	for i := 0; i < len(nums); i++ {
		start = nums[i]
		for (i+1) < len(nums) && (nums[i+1] == nums[i]+1) {
			i++
		}
		if start == nums[i] {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}

	return result
}
