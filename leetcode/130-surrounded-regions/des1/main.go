package des1

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		dfs(board, i, 0)
		dfs(board, i, col-1)
	}
	for i := 0; i < col; i++ {
		dfs(board, 0, i)
		dfs(board, row-1, i)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'A':
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, col int) {
	if board[row][col] != 'O' {
		return
	}
	colsLen := len(board[0])
	rowsLen := len(board)

	board[row][col] = 'A'
	if col < colsLen-1 {
		dfs(board, row, col+1)
	}
	if row < rowsLen-1 {
		dfs(board, row+1, col)
	}
	if col > 0 {
		dfs(board, row, col-1)
	}
	if row > 0 {
		dfs(board, row-1, col)
	}
}
