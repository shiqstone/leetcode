package code017

var m = map[byte][]string{
	'0': {},
	'1': {},
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(input string) []string {
	if input == "" {
		return nil
	}
	ret := []string{""}

	for i:=0; i<len(input); i++ {
		temp := []string{}

		for j:=0; j<len(ret); j++ {
			for k:=0; k<len(m[input[i]]); k++ {
				temp = append(temp, ret[j]+m[input[i]][k])
			}
		}
		ret = temp
	}
	return ret
}

//recuise
func letterCombinations2(input string) []string {
	if input == "" {
		return nil
	}
	ret := []string{}
	ret = addNumbers(ret, 0, "", input)
	return ret
}
func addNumbers(ret []string, i int, curstr string, input string) []string {
	if i == len(input) {
		ret = append(ret, curstr)
		return ret
	}
	candiate := m[input[i]]
	for _, char := range candiate {
		ret = addNumbers(ret, i+1, curstr+char, input)
	}
	return ret
}


