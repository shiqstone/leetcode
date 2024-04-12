package code079

func exist(board [][]byte, word string) bool {
	// Define directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m := len(board)
	n := len(board[0])

	// Define the helper function
	var backtrack func(i, j, index int) bool

	// Backtracking function
	backtrack = func(i, j, index int) bool {
		// Base case: word found
		if index == len(word) {
			return true
		}

		// Check boundaries and character match
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[index] {
			return false
		}

		// Mark current cell as visited
		temp := board[i][j]
		board[i][j] = '#'

		// Explore neighbors
		for _, dir := range directions {
			if backtrack(i+dir[0], j+dir[1], index+1) {
				return true
			}
		}

		// Backtrack by marking current cell as unvisited
		board[i][j] = temp

		return false
	}

	// Start backtracking from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
