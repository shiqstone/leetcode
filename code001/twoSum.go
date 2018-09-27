package code001

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))

	for key, val := range nums {
		if tk, ok := temp[target-val]; ok {
			return []int{tk, key}
		}
		temp[val] = key
	}
	return nil
}
