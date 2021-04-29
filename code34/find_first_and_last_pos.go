package main

func searchPos(nums []int, target int) []int {
	nlen := len(nums)
	res := []int{-1, -1}
	start := 0
	end := nlen - 1

	if nlen == 0 {
		return res
	}

	var mid int
	for start <= end {
		mid = (start + end) / 2
		if target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if start == nlen || nums[start] != target {
		return res
	} else {
		res[0] = start
	}

	start = 0
	end = nlen - 1
	for start <= end {
		mid = (start + end) / 2
		if target == nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	res[1] = end
	return res
}
