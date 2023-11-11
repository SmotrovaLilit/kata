package des2

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
	colsLen := len(board[0])
	rowsLen := len(board)
	q := &queue{}
	q.add(row, col)
	for !q.isEmpty() {
		row, col = q.pollLast()
		if board[row][col] != 'O' {
			continue
		}
		board[row][col] = 'A'
		if col < colsLen-1 && board[row][col+1] == 'O' {
			q.add(row, col+1)
		}
		if row < rowsLen-1 && board[row+1][col] == 'O' {
			q.add(row+1, col)
		}
		if col > 0 && board[row][col-1] == 'O' {
			q.add(row, col-1)
		}
		if row > 0 && board[row-1][col] == 'O' {
			q.add(row-1, col)
		}
	}

}

type node struct {
	row  int
	col  int
	next *node
}

type queue struct {
	top *node
}

func (q *queue) isEmpty() bool {
	return q.top == nil
}

func (q *queue) add(row, col int) {
	q.top = &node{
		row:  row,
		col:  col,
		next: q.top,
	}
}

func (q *queue) pollLast() (int, int) {
	if q.isEmpty() {
		panic("queue is empty")
	}
	row := q.top.row
	col := q.top.col
	q.top = q.top.next
	return row, col
}
