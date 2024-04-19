package code084

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		var h int
		if i == len(heights) {
			h = 0
		} else {
			h = heights[i]
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

func largestRectangleArea2(heights []int) int {
	heights = append([]int{0}, append(heights, 0)...)
	stack := []int{}
	stack = append(stack, 0)
	ans := 0

	for i := 1; i < len(heights); i++ {
		topIdx := stack[len(stack)-1]
		if heights[i] >= heights[topIdx] {
			stack = append(stack, i)
		} else {
			for heights[i] < heights[topIdx] {
				if len(stack)-2 < 0 {
					break
				}
				h := heights[topIdx]
				l := stack[len(stack)-2]
				r := i
				if h*(r-l-1) > ans {
					ans = h * (r - l - 1)
				}
				stack = stack[:len(stack)-1]
				topIdx = stack[len(stack)-1]
			}

			stack = append(stack, i)
		}
	}
	return ans
}
