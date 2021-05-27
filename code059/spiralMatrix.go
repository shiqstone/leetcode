package code059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	const (
		dR = iota
		dD
		dL
		dU
	)

	i, j := 0, 0
	rn, cn, rs, cs := n, n, 0, 0
	direct := dR
	v := 1
	for {
		if i == rn || j == cn || i < rs || j < cs {
			break
		}
		matrix[i][j] = v
		v++

		switch direct {
		case dR:
			if j+1 == cn {
				direct = dD
				i++
				rs++
			} else {
				j++
			}
		case dD:
			if i+1 == rn {
				direct = dL
				j--
				cn--
			} else {
				i++
			}
		case dL:
			if j == cs {
				direct = dU
				i--
				rn--
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

	return matrix
}
