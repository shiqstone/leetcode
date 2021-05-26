package code068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	lenLeft := maxWidth
	res := make([]string, 0)
	wslen := len(words)
	line := ""
	spsPos := make([]int, 0)
	for i := 0; i < wslen; i++ {
		word := words[i]
		wlen := len(word)
		if wlen > lenLeft {
			line = line[:len(line)-1]
			spLen := len(spsPos) - 1
			lenLeft += 1
			spsPos = spsPos[:spLen]
			srepLen := 0
			mrep := lenLeft
			if spLen > 0 {
				srepLen = lenLeft / spLen
				mrep = lenLeft % spLen
			}

			slen := 0
			for n, p := range spsPos {
				plen := srepLen
				sreps := strings.Repeat(" ", srepLen)
				if n < mrep {
					sreps += " "
					plen += 1
				}
				slen += plen
				line = line[:p] + sreps + line[p:]
				if n < spLen-1 {
					spsPos[n+1] += slen
				}
			}
			spsPos = make([]int, 0)
			res = append(res, line)
			line = ""
			lenLeft = maxWidth
			i--
		} else if wlen == lenLeft {
			line += word
			res = append(res, line)
			line = ""
			lenLeft = maxWidth
			spsPos = make([]int, 0)
		} else {
			line += word + " "
			spsPos = append(spsPos, len(line)-1)
			lenLeft -= wlen + 1
			if i == wslen-1 || (i < wslen-1 && len(spsPos) < 2 && len(words[i+1]) > lenLeft) {
				line += strings.Repeat(" ", lenLeft)
				res = append(res, line)
				line = ""
				lenLeft = maxWidth
				spsPos = make([]int, 0)
			} else {
				if lenLeft == 0 {
					pos := spsPos[0]
					line = line[:pos] + " " + line[pos:maxWidth-1]
					res = append(res, line)
					line = ""
					lenLeft = maxWidth
					spsPos = make([]int, 0)
				}
			}
		}

	}
	return res
}
