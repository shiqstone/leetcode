package code057

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	inserted := false

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			// If the current interval ends before the new interval starts, add it to the result
			result = append(result, interval)
		} else if interval[0] > newInterval[1] {
			// If the current interval starts after the new interval ends, and the new interval has not been inserted yet,
			// add the new interval to the result and mark it as inserted
			if !inserted {
				result = append(result, newInterval)
				inserted = true
			}
			// Add the current interval to the result
			result = append(result, interval)
		} else {
			// If there is an overlap between the current interval and the new interval, merge them
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	if !inserted {
		// If the new interval has not been inserted yet, add it to the end of the result
		result = append(result, newInterval)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
