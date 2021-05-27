package code054

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	rs := 0
	cs := 0

	res := make([]int, 0)
	i := 0
	j := 0

	const (
		dR = iota
		dD
		dL
		dU
	)

	direct := dR
	for {
		if i == m || j == n || i < rs || j < cs {
			break
		}

		res = append(res, matrix[i][j])
		matrix[i][j] = 0

		switch direct {
		case dR:
			if j+1 == n {
				direct = dD
				i++
				rs++
			} else {
				j++
			}
		case dD:
			if i+1 == m {
				direct = dL
				j--
				n--
			} else {
				i++
			}
		case dL:
			if j == cs {
				direct = dU
				i--
				m--
			} else {
				j--
			}
		case dU:
			if i == rs {
				direct = dR
				j++
				cs++
			} else {
				i--
			}
		}
	}
	return res
}
