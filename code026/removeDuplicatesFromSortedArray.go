package code026

func removeDuplicates(input []int) int {
	left, right, size := 0, 1, len(input)
	for ; right<size; right++ {
		if input[left] == input[right] {
			continue
		}
		left++
		input[left], input[right] = input[right], input[left]
	}
	return left+1
}
