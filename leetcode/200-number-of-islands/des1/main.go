package main

func dfs(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
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
				dfs(grid, r, c)
			}
		}
	}
	return numIslands
}
