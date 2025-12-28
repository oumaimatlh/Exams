package piscine

import (
	"fmt"
	"os"
	"strconv"
)

func Fprime() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if n < 0 {
		return
	}

	str := ""
	// r := n
	i := 2

	for n != 1 {

		if IsPrime(i) {
			for n%i == 0 {
				str += Itoa(i) + "*"
				n /= i
			}
		}
		i++

	}
	str = str[:len(str)-1]

	fmt.Println(str)
}
