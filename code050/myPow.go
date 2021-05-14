package code050

func myPow(x float64, n int) float64 {
	if x == 1.0 {
		return 1
	}
	if x == -1 {
		if (n & 1) != 0 {
			return -1
		} else {
			return 1
		}
	}
	if n == -2147483648 {
		return 0
	}
	res := 1.0
	if n > 0 {
		res = powRecursion(x, n)
	} else {
		res = 1 / powRecursion(x, -n)
	}
	return res
}

func powRecursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		temp := powRecursion(x, n/2)
		return temp * temp
	} else {
		temp := powRecursion(x, n/2)
		return temp * temp * x
	}
}

func myPow1(x float64, n int) float64 {
	if x == 1.0 {
		return 1
	}
	if x == -1 {
		if (n & 1) != 0 {
			return -1
		} else {
			return 1
		}
	}
	if n == -2147483648 {
		return 0
	}
	res := 1.0
	if n > 0 {
		for i := 0; i < n; i++ {
			res *= x
		}
	} else {
		for i := n; i < 0; i++ {
			res *= x
		}
		res = 1 / res
	}
	return res
}
