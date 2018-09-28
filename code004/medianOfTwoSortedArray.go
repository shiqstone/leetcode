package code004

import (
	"math"
)

func medianOfTwoSortedArray(ts1 []int, ts2 []int) float64 {
	all := make([]int, 0, len(ts1)+len(ts2))
	idx := 0
	for len(ts1) > 0 || len(ts2) > 0 {
		len1 := len(ts1)
		len2 := len(ts2)
		var t1, t2 int
		if len1 > 0 {
			t1 = ts1[0]
		}
		if len2 > 0 {
			t2 = ts2[0]
		}

		if len1 == 0 {
			all = append(all, t2)
			ts2 = ts2[1:]
			idx++
		} else if len2 == 0 {
			all = append(all, t1)
			ts1 = ts1[1:]
			idx++
		} else {
			if t1 < t2 {
				all = append(all, t1)
				ts1 = ts1[1:]
				idx++
			} else if t2 < t1 {
				all = append(all, t2)
				ts2 = ts2[1:]
				idx++
			} else {
				all = append(all, t1)
				all = append(all, t2)
				ts1 = ts1[1:]
				ts2 = ts2[1:]
				idx += 2
			}
		}
	}
	var mark int
	var median float64
	if idx%2 == 0 {
		mark = idx / 2
		median = (float64(all[mark]) + float64(all[mark-1])) / 2
	} else {
		mark = int(math.Ceil(float64(idx / 2)))
		median = float64(all[mark])
	}
	return median
}
