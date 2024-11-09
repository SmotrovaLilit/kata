package des1

func setZeroes(matrix [][]int) {
	firstRowHasZero := false
	firstColHasZero := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				continue
			}
			if i == 0 {
				firstRowHasZero = true
			}
			if j == 0 {
				firstColHasZero = true
			}
			matrix[i][0] = 0
			matrix[0][j] = 0
		}
	}

	for i := len(matrix) - 1; i > 0; i-- {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := len(matrix[0]) - 1; i > 0; i-- {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if firstRowHasZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if firstColHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
