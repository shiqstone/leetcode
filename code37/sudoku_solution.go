package code37

func solveSudoku(board [][]byte) [][]byte {
	return solve(board, 0, 0)
}

func solve(board [][]byte, i, j int) [][]byte {
	for ; i < 9; i++ {
		if j >= 9 {
			j = 0
		}
		for ; j < 9; j++ {
			if board[i][j] == '.' {
				var char byte = '1'
				for char <= '9' {
					if isValid(board, i, j, char) {
						board[i][j] = char
						res := solve(board, i, j+1)
						if res != nil {
							return res
						} else {
							board[i][j] = '.'
						}
					}
					char++
				}
				return nil
			}
		}
	}
	return board
}

func isValid(board [][]byte, row, col int, char byte) bool {
	for _, v := range board[row] {
		if v == char {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == char {
			return false
		}
	}

	var starti int = row / 3 * 3
	var startj int = col / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[starti+i][startj+j] == char {
				return false
			}
		}
	}
	return true
}
