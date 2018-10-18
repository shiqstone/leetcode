package code009

func palindromeNum(para int) bool {
	if para < 0 {
		return false
	}
	res := 0
	/*
		avoid integer overflow after reverse
		divide input to two part, compare first half with second halt reverse
	 */
	for para > res {
		res = res * 10 + para%10
		para = para/10
	}
	return res==para || res/10==para
}
