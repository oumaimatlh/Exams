package piscine

import (
	//"fmt"
	"strconv"

)

func AtoiDec(s string) int {
	n := 0
	signe := 1
	if s[0] == '-' {
		s = s[1:]
		signe = -1
	}
	for _, r := range s {
		n = n * 10 + int(r-'0')
	}
	return n * signe
}

func NotDecimal(dec string) string {

	if len(dec) == 0 {
		return "\n"
	}
	for _, r := range dec {
		if !(r >='0' && r <= '9' || r == '.' || r == '-') {
			return dec+"\n"
		}
	}
	k := ""
	for _, r := range dec {
		if r != '.'{
			k += string(r)
		}
	}

	n:= AtoiDec(k)
	res := strconv.Itoa(n)
	
	return res + "\n"


}