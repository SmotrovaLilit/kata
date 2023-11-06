package des1

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	lenI := len(matrix)
	lenJ := len(matrix[0])
	startI := 0
	startJ := 0
	stage := 0
	for lenI > 0 && lenJ > 0 && startI < lenI && startJ < lenJ {
		switch stage {
		case 0: // left to right
			for k := startJ; k < lenJ; k++ {
				res = append(res, matrix[startI][k])
			}
			startI++
			stage = 1
		case 1: // top to bottom
			for k := startI; k < lenI; k++ {
				res = append(res, matrix[k][lenJ-1])
			}
			lenJ--
			stage = 2
		case 2: // right to left
			for k := lenJ - 1; k >= startJ; k-- {
				res = append(res, matrix[lenI-1][k])
			}
			lenI--
			stage = 3
		case 3: // bottom to top
			for k := lenI - 1; k >= startI; k-- {
				res = append(res, matrix[k][startJ])
			}
			startJ++
			stage = 0
		}
	}
	return res
}
