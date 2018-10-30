package code011

func maxArea(input []int) int {
	max := 0
	area := 0
	i := 0
	j := len(input) - 1
	for i < j {
		if input[i] > input[j] {
			area = (j - i) * input[j]
			j--
		} else {
			area = (j - i) * input[i]
			i++

		}
		if area > max {
			max = area
		}
	}
	return max
}
