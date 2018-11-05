package code015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para []int
	ans  [][]int
}{
	{
		[]int{1, -1, -1, 0},
		[][]int{{-1, 0, 1}},
	},
	{
		[]int{-1, 0, 1, 2, 2, 2, 2, -1, -4},
		[][]int{
			{-4, 2, 2},
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
	{
		[]int{-1, 0, 1, 2, -1, -4},
		[][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
	{
		[]int{0, 0, 0, 0, 0},
		[][]int{{0, 0, 0}},
	},
}

func Test_threeSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := threeSum(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

func Test_threeSum2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := threeSum2(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}


func Benchmark_threeSum(t *testing.B) {
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		threeSum(tcs[0].para)
	}
}

func Benchmark_threeSum2(t *testing.B) {
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		threeSum2(tcs[0].para)
	}
}
