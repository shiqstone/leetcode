package code007

func reverseInt(para int32) int32 {
	flag := 0
	res := 0
	if para < 0 {
		flag = -1
		para = -para
	} else {
		flag = 1
	}
	for para > 0 {
		unit := para % 10
		res = res*10 + int(unit)
		para /= 10
		if (flag > 0 && res > 2147483648) || (flag < 0 && res > 2147483647) {
			res = 0
			break
		}
	}
	if flag < 0 {
		res = -res
	}
	return int32(res)
}