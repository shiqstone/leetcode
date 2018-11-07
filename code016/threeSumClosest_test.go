package code016

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct{
	lst []int
	target int
}

type tc struct {
	para param
	ans int
}

var tcs = []struct{
	para param
	ans  int
}{
	{param{
		[]int{-1, 2, 1, -4},
		1,
	},
		2,
	},
	{param{
		[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4},
		1,
	},
		0,
	},
	{param{
		[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4},
		0,
	},
		0,
	},
}


func Test_threeSumClosest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := threeSumClosest(tc.para.lst, tc.para.target)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}

//
//
//func Benchmark_threeSum(t *testing.B) {
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		threeSum(tcs[0].para)
//	}
//}
