package code012

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para  int
	ans   string
}{
	{
		3,
		"III",
	},
	{
		4,
		"IV",
	},
	{
		9,
		"IX",
	},
	{
		13,
		"XIII",
	},
	{
		29,
		"XXIX",
	},
	{
		58,
		"LVIII",
	},
	{
		1888,
		"MDCCCLXXXVIII",
	},
	{
		1994,
		"MCMXCIV",
	},
	{
		3999,
		"MMMCMXCIX",
	},
}

func Test_intToRoman(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := intToRoman(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

func BenchmarkIntToRoman(t *testing.B) {
	t.ResetTimer()
	for i:=0; i<t.N; i++{
		intToRoman(1888)
	}
}

func BenchmarkIntToRoman2(t *testing.B) {
	t.ResetTimer()
	for i:=0; i<t.N; i++{
		intToRoman2(1888)
	}
}