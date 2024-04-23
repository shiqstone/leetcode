package code088

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1    // Pointer for nums1
	p2 := n - 1    // Pointer for nums2
	p := m + n - 1 // Pointer for the merged array

	// Merge nums1 and nums2 backwards
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2, merge them into nums1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
