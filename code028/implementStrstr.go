package code028

func implementStrstr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return 0
	}
	nlen, hlen := len(needle), len(haystack)
	if nlen > hlen {
		return -1
	}
	for i:=0; i<hlen-nlen; i++ {
		if haystack[i:i+hlen] == needle {
			return i
		}
	}
	return -1
}