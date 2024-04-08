package code076

import "math"

func minWindow(s string, t string) string {
	targetFreq := make(map[byte]int)
	windowFreq := make(map[byte]int)

	// Calculate frequency of characters in string t
	for i := range t {
		targetFreq[t[i]]++
	}

	left, right := 0, 0
	minLength := math.MaxInt32
	minWindow := ""
	found := 0

	for right < len(s) {
		// Expand the window to the right
		windowFreq[s[right]]++
		if windowFreq[s[right]] <= targetFreq[s[right]] {
			found++
		}

		// Move the left pointer to minimize the window
		for found == len(t) {
			if right-left+1 < minLength {
				minLength = right - left + 1
				minWindow = s[left : right+1]
			}

			windowFreq[s[left]]--
			if windowFreq[s[left]] < targetFreq[s[left]] {
				found--
			}
			left++
		}

		right++
	}

	return minWindow
}
