package code058

func lengthOfLastWord(s string) int {
	cnt := 0
	st := false
	slen := len(s)
	for i := slen - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if st {
				break
			} else {
				continue
			}
		} else {
			if !st {
				st = true
			}
			cnt++
		}
	}

	return cnt
}
