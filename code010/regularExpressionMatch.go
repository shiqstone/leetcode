package code010

func isMatch(input string, parten string) bool {
	return checkReg(input, len(input)-1, parten, len(parten)-1)
}

//Backtracking
func checkReg(s string, i int, p string, j int) bool {
	if j == -1 {
		if i == -1 {
			return true
		} else {
			return false
		}
	}
	if j > -1 && i == -1 && p[j] != '*'{
		return false
	}
	if p[j] == '*' {
		if i > -1 && (p[j-1] == '.' || p[j-1] == s[i]) {
			if checkReg(s, i-1, p, j) {
				return true
			}
		}
		return checkReg(s, i, p, j-2)

	}
	if p[j] == '.' || p[j] == s[i] {
		return checkReg(s, i-1, p, j-1)
	}
	return false
}
