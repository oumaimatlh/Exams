package piscine

//import "strconv"

func Itoa(n int)string{
	// s := strconv.Itoa(n)
	// return  s

	s := ""
	if n == 0 {
		return  "0"
	}
	signe := true
	if n < 0 {
		signe = false
		n = -n
	}

	for n > 0 {
		r := n % 10
		s = string(rune(r+'0')) + s
		n /= 10
	}

	if !signe {
		s = "-"+s
	}
	return s
}