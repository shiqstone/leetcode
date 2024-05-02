package code063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// Initialize dp array
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Base case: if the start cell is an obstacle, return 0
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// Initialize the first row and column
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Calculate the number of unique paths for each cell
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	// Return the number of unique paths to reach the bottom-right corner
	return dp[m-1][n-1]
}
