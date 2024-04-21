package code085

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([][]int, rows)
	for i := range heights {
		heights[i] = make([]int, cols)
	}

	// Convert matrix to heights matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == 0 {
					heights[i][j] = 1
				} else {
					heights[i][j] = heights[i-1][j] + 1
				}
			}
		}
	}

	maxArea := 0
	for i := 0; i < rows; i++ {
		maxArea = max(maxArea, largestRectangleArea(heights[i]))
	}

	return maxArea
}

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	ans := 0
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, largestRectangleArea2(heights))
	}
	return ans
}

func largestRectangleArea2(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}

	ans := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	stack := []int{0}
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					l := stack[len(stack)-1]
					r := i
					h := heights[mid]
					ans = max(ans, (r-l-1)*h)
				}
			}
			stack = append(stack, i)
		}
	}

	return ans
}
