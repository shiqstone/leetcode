package code007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	para  int32
	ans   int32
}{
	{
		123,
		321,
	},
	{
		-123,
		-321,
	},
	{
		1234567899,
		0,
	},
	{
		-1234567899,
		0,
	},
}

func Test_reverseInt(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		//fmt.Printf("---%v~~~\n", tc)
		rescnt := reverseInt(tc.para)
		ast.Equal(tc.ans, rescnt, "input:%v", tc)
	}
}
