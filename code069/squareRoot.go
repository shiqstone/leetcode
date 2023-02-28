package code069

func sqrtRoot(x int) (res int) {
	if x == 0 {
		return 0
	}

	min := 1
	max := x

	res = 0
	for max-min > 1 {
		m := (max + min) / 2
		if x/m < m {
			max = m
		} else {
			min = m
		}
	}
	return min
}
