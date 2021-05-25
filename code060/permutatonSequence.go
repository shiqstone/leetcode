package code060

import "strconv"

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	res := getSub(nums, n, k)
	return res
}

func getSub(nums []int, n, k int) string {
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	res := ""
	cnt := factorial(n - 1)
	idx := (k - 1) / cnt
	s := strconv.Itoa(nums[idx])
	nums = append(nums[:idx], nums[idx+1:]...)
	k = k % cnt
	if k == 0 {
		k = cnt
	}
	res = s + getSub(nums, n-1, k)
	return res
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
