package code41

func firstMissingPositive(nums []int) int {
	slen := len(nums)
	for i := 0; i < slen; i++ {
		for nums[i] > 0 && nums[i] <= slen && nums[i] != nums[nums[i]-1] {
			swap(nums, i, nums[i]-1)
		}
	}

	for k, v := range nums {
		if v != k+1 {
			return k + 1
		}
	}
	return slen + 1
}

func swap(nums []int, i, j int) []int {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
	return nums
}
