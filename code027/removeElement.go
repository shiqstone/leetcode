package code027

func removeElement(input []int, val int) []int {
	left, right := 0, len(input)-1
	for {
		for left<len(input) && input[left] != val {
			left++
		}
		for right>=0 && input[right] == val {
			right--
		}

		if left >= right {
			break;
		}
		input[left], input[right] = input[right], input[left]
	}
	return input[:left]
}
