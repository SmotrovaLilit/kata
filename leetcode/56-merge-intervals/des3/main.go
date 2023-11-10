package des1

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	lastIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[lastIndex][1] < intervals[i][0] { // not overlapped
			lastIndex++
			continue
		}
		intervals[lastIndex][1] = max(intervals[lastIndex][1], intervals[i][1])
		intervals = append(intervals[:i], intervals[i+1:]...)
		i--
	}
	return intervals
}
