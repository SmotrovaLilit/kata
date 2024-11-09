package des2

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		bfs(board, i, 0)
		bfs(board, i, col-1)
	}
	for i := 0; i < col; i++ {
		bfs(board, 0, i)
		bfs(board, row-1, i)
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

func bfs(board [][]byte, row, col int) {
	colsLen := len(board[0])
	rowsLen := len(board)
	q := &queue{}
	q.add(row, col)
	for !q.isEmpty() {
		row, col = q.pollFirst()
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
	head *node
	tail *node
}

func (q *queue) isEmpty() bool {
	return q.head == nil
}

func (q *queue) add(row, col int) {
	if q.head == nil {
		q.head = &node{
			row: row,
			col: col,
		}
		q.tail = q.head
		return
	}
	q.tail.next = &node{
		row: row,
		col: col,
	}
	q.tail = q.tail.next
}

func (q *queue) pollFirst() (int, int) {
	if q.isEmpty() {
		panic("queue is empty")
	}
	row := q.head.row
	col := q.head.col
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return row, col
}
