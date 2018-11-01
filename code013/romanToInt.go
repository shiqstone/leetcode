package code013

func romanToInt(input string) int {
	/*
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000

	input 1~3999
	*/

	res := 0
	m := map[byte]int {
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}

	last := 0
	for i := len(input)-1; i>=0; i-- {
		temp := m[input[i]]
		sign := 1
		if temp < last {
			sign = -1
		}
		res += sign * temp

		last = temp
	}
	return res
}

