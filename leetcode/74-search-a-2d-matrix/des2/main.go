package des1

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	start := 0
	end := m*n - 1
	var virtualMid, virtualMidValue int
	for start <= end {
		virtualMid = start + (end-start)/2
		virtualMidValue = matrix[virtualMid/n][virtualMid%n]
		if virtualMidValue == target {
			return true
		}

		if target < virtualMidValue {
			end = virtualMid - 1
			continue
		}
		start = virtualMid + 1
	}
	return false
}
