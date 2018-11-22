package code018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	arr    []int
	target int
}

var tcs = []struct {
	p   para
	ans [][]int
}{
	{
		para{
			[]int{1, 0, -1, 0, -2, 2}, 0,
		},
		[][]int{
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
			[]int{-1, 0, 0, 1},
		},
	},
	{
		para{
			[]int{-1, -1, -2, -2, 1, 1, 2, 2, 0, 0, 0, 0}, 0,
		},
		[][]int{
			[]int{-2, -2, 2, 2},
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
			[]int{-2, 0, 1, 1},
			[]int{-1, -1, 0, 2},
			[]int{-1, -1, 1, 1},
			[]int{-1, 0, 0, 1},
			[]int{0, 0, 0, 0},
		},
	},
	{
		para{
			[]int{0, -1, 0, 1, -2, -5, 3, 5, 0}, 6,
		},
		[][]int{
			[]int{-2, 0, 3, 5},
			[]int{0, 0, 1, 5},
		},
	},
	{
		para{
			[]int{0, 0, 0, 0}, 0,
		},
		[][]int{
			[]int{0, 0, 0, 0},
		},
	},
}

func Test_fourSum(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := fourSum(tc.p.arr, tc.p.target)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

//func Test_fourSum2(t *testing.T) {
//	ast := assert.New(t)
//
//	for _, tc := range tcs {
//		//fmt.Printf("---%v~~~\n", tc)
//		rescnt := fourSum2(tc.para)
//		ast.Equal(tc.ans, rescnt, "input:%v", tc)
//	}
//}

func Benchmark_fourSum(t *testing.B) {
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fourSum(tcs[0].p.arr, tcs[0].p.target)
	}
}

//func Benchmark_fourSum2(t *testing.B) {
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		fourSum2(tcs[0].para)
//	}
//}
