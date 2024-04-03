package code072

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	// Initialize a 2D array for dynamic programming
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Perform dynamic programming
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insertion := dp[i][j-1] + 1
				deletion := dp[i-1][j] + 1
				replacement := dp[i-1][j-1] + 1
				dp[i][j] = min(insertion, min(deletion, replacement))
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
