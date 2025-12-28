package piscine

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	signe := 1
	if s[0] == '-' {
		signe = -1
		s = s[1:]
	}

	n := 0
	for _, r := range s {
		if !(r >= '0' && r <= '9') {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return signe * n
}

func AddPrimeSum() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	n := Atoi(os.Args[1])
	if n <= 0 {
		fmt.Println(0)
		return
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
