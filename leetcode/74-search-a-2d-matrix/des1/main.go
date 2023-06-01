package des1

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	startM := 0
	endM := m - 1
	foundRow := 0
	for startM <= endM {
		i := startM + (endM-startM)/2
		if matrix[i][n-1] == target {
			return true
		}
		if matrix[i][0] == target {
			return true
		}
		if target < matrix[i][n-1] && target > matrix[i][0] {
			foundRow = i
			break
		}
		if target < matrix[i][n-1] {
			endM = i - 1
			continue
		}
		startM = i + 1
	}
	startN := 1
	endN := n - 2
	for startN <= endN {
		i := startN + (endN-startN)/2
		if matrix[foundRow][i] == target {
			return true
		}
		if target < matrix[foundRow][i] {
			endN = i - 1
			continue
		}
		startN = i + 1
	}
	return false
}
