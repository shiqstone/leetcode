package code049

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	rmap := make(map[string]([]string))

	for _, str := range strs {
		num := make([]byte, 26)
		for _, v := range str {
			num[v-'a']++
		}
		tmp := string(num)
		rmap[tmp] = append(rmap[tmp], str)
	}
	res := make([][]string, 0)
	for _, v := range rmap {
		res = append(res, v)
	}
	return res
}

func groupAnagrams1(strs []string) [][]string {
	rmap := make(map[string]([]string))

	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		tmp := strings.Join(s, "")

		if v, ok := rmap[tmp]; ok {
			v := append(v, str)
			rmap[tmp] = v
		} else {
			rmap[tmp] = []string{str}
		}
	}
	res := make([][]string, 0)
	for _, v := range rmap {
		res = append(res, v)
	}
	return res
}
