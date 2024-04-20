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
