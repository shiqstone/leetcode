package code042

func trap(height []int) int {
	ilen := len(height)
	left := make([]int, ilen)
	right := make([]int, ilen)
	for i := 0; i < ilen; i++ {
		left[i] = 0
		right[i] = 0
	}

	for i := 0; i < ilen-1; i++ {
		if height[i] > left[i+1] {
			left[i+1] = height[i]
		}
		if i > 0 && left[i-1] > left[i] {
			left[i] = left[i-1]
		}
	}
	for i := ilen - 1; i > 0; i-- {
		if height[i] > right[i-1] {
			right[i-1] = height[i]
		}
		if i < ilen-1 && right[i+1] > right[i] {
			right[i] = right[i+1]
		}
	}

	res := 0
	for i := 0; i < ilen; i++ {
		top := 0
		if left[i] > right[i] {
			top = right[i]
		} else {
			top = left[i]
		}
		if top >= height[i] {
			res += top - height[i]
		}
	}
	return res
}
