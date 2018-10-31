package code012

import "strings"

func intToRoman(input int) string {
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
	var roman string
	dict := [][]string {
		{"I", "V"},
		{"X", "L"},
		{"C", "D"},
		{"M"},
	}
	d := 0
	var temp []string
	for input > 0 {
		dec := input % 10
		switch {
		case dec < 4 :
			temp = append(temp, strings.Repeat(dict[d][0], dec))
		case dec == 4 :
			temp = append(temp, dict[d][0]+dict[d][1])
		case dec <9 :
			hdec := dec % 5
			temp = append(temp, dict[d][1]+strings.Repeat(dict[d][0], hdec))
		case dec == 9:
			temp = append(temp, dict[d][0]+dict[(d+1)][0])
		}
		input /= 10
		d++
	}
	for i:=len(temp)-1; i>=0; i--{
		roman += temp[i]
	}
	return roman
}

func intToRoman2(num int) string {

	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}
