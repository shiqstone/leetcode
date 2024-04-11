package code078

func subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	// Define the helper function
	var backtrack func(start int)

	// Backtracking function
	backtrack = func(start int) {
		// Add current subset to result
		temp := make([]int, len(subset))
		copy(temp, subset)
		result = append(result, temp)

		// Iterate over numbers from start to end of nums
		for i := start; i < len(nums); i++ {
			// Add current number to subset
			subset = append(subset, nums[i])
			// Explore next index
			backtrack(i + 1)
			// Backtrack by removing current number
			subset = subset[:len(subset)-1]
		}
	}

	// Start backtracking from index 0
	backtrack(0)

	return result
}
