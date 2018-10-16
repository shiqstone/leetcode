package code006

import "bytes"

func convert(strline string, rows int) string {
	if rows == 1 || len(strline) < rows {
		return strline
	}
	res := bytes.Buffer{}

	//p pace
	p := rows*2 - 2
	lens := len(strline)

	//line 1    n * p
	for i := 0; i < lens; i += p {
		res.WriteByte(strline[i])
	}
	//mid line r，1×p-r，1×p+r，2×p-r，2×p+r，.
	for r := 1; r <= rows-2; r++ {
		res.WriteByte(strline[r])
		for k := p; k-r < lens; k += p {
			res.WriteByte(strline[k-r])
			if k+r < lens {
				res.WriteByte(strline[k+r])
			}
		}
	}
	//line last row - 1 + n * p
	for i := rows - 1; i < lens; i += p {
		res.WriteByte(strline[i])
	}
	return res.String()
}

func convert2(strline string, rows int) string {
	/*
		access by line
		line 0  char pos: (2*rows-2)*n;
		line mid char pos: (2*rows-2)*n and (n+1)*(2*rows-2)-i
		line rows-1 char pos: (2*rows-2)*n + line ;
	 */
	if rows == 1 || len(strline) < rows {
		return strline
	}
	res := bytes.Buffer{}

	//p pace
	p := rows*2 - 2
	lens := len(strline)

	for r := 0; r < rows; r++ {
		for k := 0; k+r < lens; k += p {
			res.WriteByte(strline[k+r])
			if r != 0 && r != rows-1 && k+p-r < lens {
				res.WriteByte(strline[k+p-r])
			}
		}
	}
	return res.String()
}
