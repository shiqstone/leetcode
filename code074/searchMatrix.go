package code074

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left := 0
	right := n - 1
	row := 0
	col := right

	for row < m {
		if matrix[row][col] < target {
			row += 1
			left = row * n
			right = left + n - 1
			continue
		} else {
			for left <= right {
				mid := (left + right) / 2
				col = mid - row*n
				element := matrix[row][col]

				if element == target {
					return true
				} else if element < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
				continue
			}
			break
		}
	}

	return false
}
