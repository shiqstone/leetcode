package code41

func firstMissingPositive(nums []int) int {
	slen := len(nums)
	for i := 0; i < slen; i++ {
		for nums[i] > 0 && nums[i] <= slen && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < slen; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return slen + 1
}
