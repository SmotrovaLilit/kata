package main

func bfs(grid [][]byte, r, c int) {
	grid[r][c] = '0'
	queue := make([][2]int, 0, len(grid))
	queue = append(queue, [2]int{r, c})
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		row := current[0]
		col := current[1]
		if row-1 >= 0 && grid[row-1][col] == '1' {
			queue = append(queue, [2]int{row - 1, col})
			grid[row-1][col] = '0'

		}
		if row+1 < len(grid) && grid[row+1][col] == '1' {
			queue = append(queue, [2]int{row + 1, col})
			grid[row+1][col] = '0'

		}
		if col-1 >= 0 && grid[row][col-1] == '1' {
			queue = append(queue, [2]int{row, col - 1})
			grid[row][col-1] = '0'
		}
		if col+1 < len(grid[0]) && grid[row][col+1] == '1' {
			queue = append(queue, [2]int{row, col + 1})
			grid[row][col+1] = '0'
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	numIslands := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' { // found root node
				numIslands++
				bfs(grid, r, c)
			}
		}
	}
	return numIslands
}
