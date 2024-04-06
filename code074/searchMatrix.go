package code074

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	left := 0
	right := m*n - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / n
		col := mid % n
		element := matrix[row][col]

		if element == target {
			return true
		} else if element < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
