package des1

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	i, j := 0, 1
	for i < len(intervals) {
		interval1 := intervals[i]
		founded := false
		for j < len(intervals) {
			interval2 := intervals[j]
			nonOverlappingIntervals := makeNonOverlappingIntervals(interval1, interval2)
			if len(nonOverlappingIntervals) != 2 { //  overlapping
				founded = true
				intervals[i] = nonOverlappingIntervals[0]
				intervals = append(intervals[:j], intervals[j+1:]...) // remove jth element
				break
			}
			j++
		}
		if founded { // start from beginning
			i = 0
			j = 1
			continue
		}
		i++
		j = i + 1
	}
	return intervals
}

func isOverlapped(a []int, b []int) bool {
	return a[0] <= b[1] && b[0] <= a[1]
}

func makeNonOverlappingIntervals(a, b []int) [][]int {
	if isOverlapped(a, b) {
		return [][]int{{min(a[0], b[0]), max(a[1], b[1])}}
	}
	return [][]int{a, b}
}
