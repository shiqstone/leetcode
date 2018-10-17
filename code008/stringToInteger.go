package code008

import (
	"math"
	"strings"
)

func myAtoi(str string) int32 {
	str = strings.TrimSpace(str)
	lens := len(str)
	flag := 1
	res := 0
	if str[0] == '-' {
		flag = -1
		str = str[1:]
		lens -= 1
	} else if str[0] == '+' {
		str = str[1:]
		lens -= 1
	}

	for i := 0; i < lens; i++ {
		charInt := int(str[i]) - 48
		if charInt < 0 || charInt > 9 {
			break
		}
		res = res*10 + charInt
		if flag > 0 && res > math.MaxInt32 {
			res = math.MaxInt32
			break
		} else if flag < 0 && res > -math.MinInt32 {
			res = math.MinInt32
			break
		}
	}
	res *= flag
	return int32(res)
}
