package code080

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
