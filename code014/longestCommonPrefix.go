package code014

func commonPrefix(input []string) string {
	prefix := ""
	lens := len(input[0])
	for _, str := range input {
		if len(str) < lens {
			lens = len(str)
		}
	}
	flag := true
	for i:=0; i<lens; i++ {
		for _, str := range input {
			if input[0][i:i+1] != str[i:i+1] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		prefix += input[0][i:i+1]
	}
	return prefix
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	short := strs[0]
	for _, s := range strs {
		if len(short) > len(s) {
			short = s
		}
	}

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}


func commonPrefix3(input []string) string {
	if len(input) == 0 {
		return ""
	}
	prefix := ""

	for i:=0; i< len(input[0]); i++ {
		for j:=1; j<len(input); j++ {
			if i>=len(input[j]) || input[0][i:i+1] != input[j][i:i+1] {
				return prefix
			}
		}
		prefix += input[0][i:i+1]
	}
	return prefix
}