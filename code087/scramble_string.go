package code087

import "fmt"

func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)
	return isScrambleRecursive(s1, s2, 0, 0, len(s1), memo)
}

func isScrambleRecursive(s1, s2 string, i, j, length int, memo map[string]bool) bool {
	key := fmt.Sprintf("%s|%s|%d|%d|%d", s1, s2, i, j, length)
	if val, ok := memo[key]; ok {
		return val
	}

	// Base case
	if s1[i:i+length] == s2[j:j+length] {
		return true
	}

	// Check if the frequency of characters in both substrings is the same
	if !isEqualFrequency(s1[i:i+length], s2[j:j+length]) {
		return false
	}

	// Check all possible lengths of the left substring
	for k := 1; k < length; k++ {
		if (isScrambleRecursive(s1, s2, i, j, k, memo) && isScrambleRecursive(s1, s2, i+k, j+k, length-k, memo)) ||
			(isScrambleRecursive(s1, s2, i, j+length-k, k, memo) && isScrambleRecursive(s1, s2, i+k, j, length-k, memo)) {
			memo[key] = true
			return true
		}
	}

	memo[key] = false
	return false
}

func isEqualFrequency(s1, s2 string) bool {
	freq := make(map[byte]int)
	for i := range s1 {
		freq[s1[i]]++
		freq[s2[i]]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}
