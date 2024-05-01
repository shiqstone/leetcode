package code062

func uniquePaths(m int, n int) int {
	// Initialize dp array
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: there is only one way to reach the cells in the first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Calculate the number of unique paths for each cell
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// Return the number of unique paths to reach the bottom-right corner
	return dp[m-1][n-1]
}
