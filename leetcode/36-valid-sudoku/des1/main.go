package des1

import "strconv"

func isValidSudoku(board [][]byte) bool {
	seenInBox := [10]byte{}
	for i := 0; i < 9; i++ {
		seenInRow := [10]byte{}
		seenInCol := [10]byte{}
		for j := 0; j < 9; j++ {
			// Rows
			if board[i][j] != '.' {
				v, err := strconv.Atoi(string(board[i][j]))
				if err != nil {
					panic(err)
				}
				seenInRow[v]++
				if seenInRow[v] > 1 {
					return false
				}
			}
			// Cols
			if board[j][i] != '.' {
				v, err := strconv.Atoi(string(board[j][i]))
				if err != nil {
					panic(err)
				}
				seenInCol[v]++
				if seenInCol[v] > 1 {
					return false
				}
			}
			// in box
			if j%3 == 0 && i%3 == 0 {
				seenInBox = [10]byte{}
				for k := i; k < i+3; k++ {
					for l := j; l < j+3; l++ {
						if board[k][l] == '.' {
							continue
						}
						v, err := strconv.Atoi(string(board[k][l]))
						if err != nil {
							panic(err)
						}
						seenInBox[v]++
						if seenInBox[v] > 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
