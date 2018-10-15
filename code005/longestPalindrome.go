package code005

import (
	"math"
	"strings"
)

func longestPalindrome(s string) string {
	//use mancher algorithm
	s = "#" + strings.Join(strings.Split(s, ""), "#") + "#"
	lens := len(s)
	rl := make([]int, lens)
	maxRight := 0
	pos := 0
	maxLen := 0
	maxStr := ""
	for i := 0; i < lens; i++ {
		if i < maxRight {
			rl[i] = int(math.Min(float64(rl[2*pos-i]), float64(maxRight-i)))
		} else {
			rl[i] = 1
		}
		//try extends bound
		for i-rl[i] >= 0 && i+rl[i] < lens && s[i-rl[i]] == s[i+rl[i]] {
			rl[i] += 1
		}
		//update maxRight & pos
		if rl[i]+i-1 > maxRight {
			maxRight = rl[i] + i - 1
			pos = i
		}
		//update longest
		if rl[i] > maxLen {
			maxLen = rl[i]
			maxStr = s[(i - rl[i] + 1):maxRight]
		}
		//fmt.Println(maxLen)
	}
	maxStr = strings.Replace(maxStr, "#", "", -1)
	return maxStr
}
