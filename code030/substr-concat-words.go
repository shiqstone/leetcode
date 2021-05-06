package code30

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)

	wcnt := len(words)
	wlen := len(words[0])
	slen := len(s)

	for i := 0; i < wlen; i++ {
		for j := i; j < (slen - wcnt*wlen + 1); j = j + wlen {
			n := 0
			tmp := make([]string, wcnt)
			copy(tmp, words)
			for {
				if n >= wcnt {
					break
				}
				substr := s[j+n*wlen : j+(n+1)*wlen]

				match := false
				for k := 0; k < len(tmp); k++ {
					word := tmp[k]
					if word == substr {
						match = true
						tmp = append(tmp[:k], tmp[k+1:]...)
						k--
						break
					}
				}

				if !match {
					j = j + n*wlen
					n = 0
					break
				}
				n++
			}
			if n == wcnt {
				res = append(res, j)
			}
		}
	}
	return res
}
