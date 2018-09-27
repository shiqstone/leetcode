package code003

func LongestSubstring(s string) int {
	chars := []byte(s)
	count := 0
	max := 0
	tchars := make(map[byte]int)
	for _, char := range chars {
		if len(tchars) == 0 {
			tchars[char] = 1
			count++
		} else if _, ok := tchars[char]; ok {
			count = 0
			tchars = make(map[byte]int)
		} else {
			tchars[char] = 1
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}