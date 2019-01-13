package code029

import "math"

func devide(m, n int) int {
	if n == 0 {
		return math.MaxInt32
	}

	ms, mabs := getSign(m)
	ns, nabs := getSign(n)

	res,_ := d(mabs, nabs, 1)

	if ms != ns {
		res = res - res -res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func getSign(m int) (int, int) {
	sign := 1
	abs := m
	if m < 0 {
		sign = -1
		abs = m - m - m
	}
	return sign, abs
}

func d(m, n, cnt int) (res, remain int) {
	if (m < n){
		return 0, m
	} else if (n<=m && m < n+n){
		return cnt, m-n
	}else{
		res, remain = d(m, n+n, cnt+cnt)
		if remain >= n {
			return res+cnt, remain -n
		}
		return
	}
}