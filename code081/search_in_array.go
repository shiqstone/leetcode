package code081

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return true
		}

		// Determine which side of the array is sorted
		if nums[left] < nums[mid] { // Left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // Handle duplicates by moving the left pointer
			left++
		}
	}

	return false
}
