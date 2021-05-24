package code066

func plusOne(digits []int) []int {
	carry := 0
	res := make([]int, 0)
	nlen := len(digits)
	for i := nlen - 1; i >= 0; i-- {
		v := digits[i] + carry
		if i == nlen-1 {
			v += 1
		}
		if v <= 9 {
			digits[i] = v
			res = digits
			break
		} else {
			carry = 1
			digits[i] = 0
			if i == 0 {
				res = append(res, 1)
				res = append(res, digits...)
			}
		}
	}
	return res
}
