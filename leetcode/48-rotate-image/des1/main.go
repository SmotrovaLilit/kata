package des1

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		rotateLayer(matrix, i)
	}
}

func rotateLayer(matrix [][]int, layerNumber int) {
	lastIndex := len(matrix) - layerNumber - 1
	var temp, row, col int
	for j := layerNumber; j < lastIndex; j++ {
		temp = matrix[layerNumber][j]
		row = layerNumber
		col = j
		for l := 0; l < 4; l++ { // to rotate on 90 degrees we need to do 4 swaps
			row, col = nextPointToRotate(row, col, layerNumber, lastIndex)
			matrix[row][col], temp = temp, matrix[row][col]
		}
	}
}

func nextPointToRotate(row int, col int, firstIndex, lastIndex int) (int, int) {
	if row == firstIndex && col < lastIndex { // to right and down
		return col, lastIndex
	}
	if col == lastIndex && row < lastIndex { // to down and left
		return lastIndex, firstIndex + lastIndex - row
	}
	if row == lastIndex && col > firstIndex { // to left and up
		return col, firstIndex
	}
	if col == firstIndex && row > firstIndex { // to up and right
		return firstIndex, firstIndex + lastIndex - row
	}
	panic("unreachable")
}
