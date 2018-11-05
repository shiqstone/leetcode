package code015

import "sort"

func threeSum(nums []int) [][]int {
	// 排序后，可以按规律查找
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0:
				// 较小的 l 需要变大
				l++
			case s > 0:
				// 较大的 r 需要变小
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 为避免重复添加，l 和 r 都需要移动到不同的元素上。
				l, r = next(nums, l, r)
			}
		}
	}

	return res
}

func threeSum2(input []int) [][]int {
	sort.Ints(input)

	tres := make([][]int, 0)

	var mid, right int

	for left := 0; left < len(input) && input[left] <= 0; left++ {
		mid = left + 1
		right = len(input) - 1
		tmp := 0 - input[left]
		if left > 0 && input[left] == input[left-1] {
			continue
		}
		for mid < right {
			if input[mid]+input[right] == tmp {
				tmid := input[mid]
				tright := input[right]
				tres = append(tres, []int{input[left], input[mid], input[right]})

				for ;mid < right && input[mid] == tmid;mid++ {
				}
				for ;mid < right && input[right] == tright;right-- {
				}
			} else if input[mid]+input[right] < tmp {
				mid++
			} else {
				right--
			}
		}
	}
	return tres
}


func next(nums []int, l, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}
