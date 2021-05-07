package code043

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)

	res := "0"
	for i := len1 - 1; i >= 0; i-- {
		dn1 := int(num1[i] - '0')
		dres := "0"
		ppi := len1 - i - 1
		for j := len2 - 1; j >= 0; j-- {
			dn2 := int(num2[j] - '0')
			tmp := dn1 * dn2
			tmpstr := strconv.Itoa(tmp)
			ppj := len2 - j - 1
			if ppj > 0 {
				for k := 0; k < ppj; k++ {
					tmpstr += "0"
				}
				// tmpstr += strings.Repeat("0", ppj)
			}

			dres = sumStr(dres, tmpstr)
		}
		if ppi > 0 {
			for k := 0; k < ppi; k++ {
				dres += "0"
			}
		}
		res = sumStr(res, dres)
	}
	return res
}

func add(a, b string) string {
	alen := len(a)
	blen := len(b)

	carry := 0
	if alen >= blen {
		for i := blen - 1; i >= 0; i-- {
			j := alen - (blen - i)
			bi, _ := strconv.Atoi(b[i : i+1])
			ai, _ := strconv.Atoi(a[j : j+1])
			tmp := bi + ai + carry
			dd := tmp % 10
			carry = tmp / 10
			a = a[:j] + strconv.Itoa(dd) + a[j+1:]
		}
		if carry > 0 {
			if alen == blen {
				a = "1" + a
			} else {
				j := alen - blen
				a = add(a[:j], "1") + a[j:]
			}
		}
		return a
	} else {
		for i := alen - 1; i >= 0; i-- {
			j := blen - (alen - i)
			ai, _ := strconv.Atoi(a[i : i+1])
			bi, _ := strconv.Atoi(b[j : j+1])
			tmp := bi + ai + carry
			dd := tmp % 10
			carry = tmp / 10
			b = b[:j] + strconv.Itoa(dd) + b[j+1:]
		}
		if carry > 0 {
			if alen == blen {
				b = "1" + b
			} else {
				j := blen - alen
				b = add(b[:j], "1") + b[j:]
			}
		}
		return b
	}
}

func sumStr(a, b string) string {
	carry := 0
	ai := len(a) - 1
	bi := len(b) - 1
	res := []byte{}
	for ai >= 0 || bi >= 0 {
		an := 0
		bn := 0
		if ai >= 0 {
			an = int(a[ai] - '0')
		}
		if bi >= 0 {
			bn = int(b[bi] - '0')
		}
		sum := an + bn + carry
		carry = sum / 10
		ans := sum % 10
		// res = strconv.Itoa(ans) + res
		res = append([]byte{byte('0' + ans)}, res...)
		ai--
		bi--
	}
	if carry > 0 {
		// res = "1" + res
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}
