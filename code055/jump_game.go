package code055

func canJump(nums []int) bool {
	maxReach := 0 // Furthest index reachable from the current position
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false // Cannot reach the current index
		}
		maxReach = max(maxReach, i+nums[i]) // Update furthest reachable index
		if maxReach >= len(nums)-1 {
			return true // Successfully reached the last index
		}
	}
	return false // Failed to reach the last index
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
