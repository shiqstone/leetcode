package code053

func maxSubArray(nums []int) int {
	dp := nums[0]
	max := nums[0]
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if dp < 0 {
			dp = v
		} else {
			dp += v
		}
		if dp > max {
			max = dp
		}
	}
	return max
}
