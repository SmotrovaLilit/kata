package des1

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if merged[len(merged)-1][1] < intervals[i][0] { // not overlapped
			merged = append(merged, intervals[i])
			continue
		}
		merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
	}
	return merged
}
