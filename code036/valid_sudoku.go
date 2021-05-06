package main

func isValidSudoku(board [][]byte) bool {
	row := make(map[byte][]int)
	col := make(map[byte][]int)
	sub := make(map[byte][]int)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			char := board[i][j]
			if char == '.' {
				continue
			}
			if rows, ok := row[char]; ok {
				rf := false
				for _, v := range rows {
					if v == i {
						rf = true
					}
				}
				if rf {
					return false
				} else {
					row[char] = append(row[char], i)
				}
			} else {
				row[char] = []int{i}
			}

			if cols, ok := col[char]; ok {
				rf := false
				for _, v := range cols {
					if v == j {
						rf = true
					}
				}
				if rf {
					return false
				} else {
					col[char] = append(col[char], j)
				}
			} else {
				col[char] = []int{j}
			}

			var si int = i / 3
			var sj int = j / 3
			idx := [][]int{
				[]int{0, 1, 2},
				[]int{3, 4, 5},
				[]int{6, 7, 8},
			}
			sv := idx[si][sj]
			if subs, ok := sub[char]; ok {
				rf := false
				for _, v := range subs {
					if v == sv {
						rf = true
					}
				}
				if rf {
					return false
				} else {
					sub[char] = append(sub[char], sv)
				}
			} else {
				sub[char] = append(sub[char], sv)
			}
		}
	}
	return true
}
