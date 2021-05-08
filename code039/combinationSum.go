package code039

func combinationSum(candidates []int, target int) [][]int {
	list := make([][]int, 0)
	tmp := make([]int, 0)
	list = backtrace(list, tmp, candidates, target, 0)
	return list
}

func backtrace(list [][]int, tmp []int, candidates []int, target int, start int) [][]int {
	if target < 0 {
		return list
	} else if target == 0 {
		slist := make([]int, len(tmp))
		copy(slist, tmp)
		list = append(list, slist)
		return list
	} else {
		for i := start; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			list = backtrace(list, tmp, candidates, target-candidates[i], i)

			//find a solution or target < 0, remove current num and keep try
			tmp = tmp[:len(tmp)-1]
		}
	}
	return list
}
