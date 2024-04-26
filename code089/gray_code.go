package code089

func grayCode(n int) []int {
	result := make([]int, 1<<uint(n)) // Initialize result slice with the required length
	for i := 0; i < len(result); i++ {
		result[i] = i ^ (i >> 1) // Generate Gray code using bitwise XOR
	}
	return result
}
