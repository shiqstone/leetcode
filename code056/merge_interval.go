package code056

import "sort"

func merge(intervals [][]int) [][]int {
	// Sort intervals based on their start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	for _, interval := range intervals {
		if len(result) == 0 || interval[0] > result[len(result)-1][1] {
			// If the current interval does not overlap with the last interval in the result array,
			// add it to the result array
			result = append(result, interval)
		} else {
			// If the current interval overlaps with the last interval in the result array,
			// merge them by updating the end time of the last interval
			result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
		}
	}
	return result
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
