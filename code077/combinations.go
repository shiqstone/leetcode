package code077

func combine(n int, k int) [][]int {
	var result [][]int
	var combination []int

	// Define the helper function
	var backtrack func(start int)

	// Backtracking function
	backtrack = func(start int) {
		// Base case: combination size reaches k
		if len(combination) == k {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		// Iterate over numbers from start to n
		for i := start; i <= n; i++ {
			// Add current number to combination
			combination = append(combination, i)
			// Explore next index
			backtrack(i + 1)
			// Backtrack by removing current number
			combination = combination[:len(combination)-1]
		}
	}

	// Start backtracking from index 1
	backtrack(1)

	return result
}
