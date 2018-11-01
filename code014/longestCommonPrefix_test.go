package code014

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para  []string
	ans   string
}{
	{
		[]string{"abcdd", "abcde", "ab"},
		"ab",
	},
	{
		[]string{"abcdd", "abcde"},
		"abcd",
	},
	{
		[]string{"dd", "abcde"},
		"",
	},
}

/*
BenchmarkCommonPrefix-4         20000000                85.8 ns/op             2 B/op          1 allocs/op
BenchmarkCommonPrefix2-4        100000000               15.6 ns/op             0 B/op          0 allocs/op
BenchmarkCommonPrefix3-4        20000000                75.1 ns/op             2 B/op          1 allocs/op
 */
func Test_commonPrefix(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := commonPrefix(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

func Test_commonPrefix3(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := commonPrefix3(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

func BenchmarkCommonPrefix(t *testing.B) {
	t.ResetTimer()
	for i:=0; i<t.N; i++{
		commonPrefix([]string{"abcdd", "abcde", "ab"})
		commonPrefix([]string{"abcdd", "abcde"})
	}
}

func BenchmarkCommonPrefix2(t *testing.B) {
	t.ResetTimer()
	for i:=0; i<t.N; i++{
		longestCommonPrefix([]string{"abcdd", "abcde", "ab"})
		longestCommonPrefix([]string{"abcdd", "abcde"})
	}
}

func BenchmarkCommonPrefix3(t *testing.B) {
	t.ResetTimer()
	for i:=0; i<t.N; i++{
		commonPrefix3([]string{"abcdd", "abcde", "ab"})
		commonPrefix3([]string{"abcdd", "abcde"})
	}
}