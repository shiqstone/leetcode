package main

func nextPermutation(nums []int) []int {
	nlen := len(nums)
	i := nlen - 2
	// for {
	// 	if i < 0 || nums[i+1] > nums[i] {
	// 		break
	// 	}
	// 	i--
	// }
	for ; i >= 0 && nums[i+1] <= nums[i]; i-- {
		//找到第一个不再递增的位置
	}

	if i < 0 {
		res := reverse(nums, 0)
		return res
	}

	j := nlen - 1
	// for {
	// 	if j < 0 || nums[j] > nums[i] {
	// 		break
	// 	}
	// 	j--
	// }
	for ; j >= 0 && nums[j] <= nums[i]; j-- {
		//找到刚好大于 nums[i]的位置
	}

	nums = swap(nums, i, j)

	nums = reverse(nums, i+1)

	return nums
}

func swap(nums []int, i, j int) []int {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
	return nums
}

func reverse(nums []int, start int) []int {
	for i, j := start, (len(nums) - 1); i < j; i, j = i+1, j-1 {
		nums = swap(nums, i, j)
	}
	return nums
}
