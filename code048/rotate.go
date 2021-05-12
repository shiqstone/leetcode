package code048

func rotate(matrix [][]int) [][]int {
	n := len(matrix)

	//notice i range 0~n/2 and j range i~(n-i-1)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
	return matrix
}
