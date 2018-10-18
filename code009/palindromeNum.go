package code009

func palindromeNum(para int) bool {
	if para < 0 {
		return false
	}
	temp := para
	res := 0
	for temp > 0 {
		res = res * 10 + temp%10
		temp = temp/10
	}
	return res==para
}
