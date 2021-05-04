package code38

import "strconv"

func countAndSay(n int) string {
	start := "1"
	for i := 1; i < n; i++ {
		slen := len(start)
		res := ""
		cnt := 1
		for j := 0; j <= slen; j++ {
			if j+2 > slen {
				res += strconv.Itoa(cnt) + string(start[j])
				break
			}
			if start[j] != start[j+1] {
				res += strconv.Itoa(cnt) + string(start[j])
				cnt = 1
			} else {
				cnt++
			}
		}
		start = res
	}
	return start
}

func countAndSay2(n int) string {
	res := "1"
	for n > 1 {
		temp := ""
		for j := 0; j < len(res); j++ {
			num := getRepeat(res[j:])
			temp = temp + strconv.Itoa(num) + string(res[j])
			j = j + num - 1
		}
		n--
		res = temp
	}
	return res
}

func getRepeat(sub string) int {
	count := 1
	char := sub[0]
	for i := 1; i < len(sub); i++ {
		if char == sub[i] {
			count++
		} else {
			break
		}
	}
	return count
}
