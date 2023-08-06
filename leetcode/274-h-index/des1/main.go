package des1

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	i := 0
	for (i < len(citations)) && (citations[i] > i) {
		i++
	}
	return i
}
