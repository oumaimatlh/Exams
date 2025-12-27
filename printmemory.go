package main

import (
	//"fmt"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	s := ""
	count := 0
	for i, r := range arr {
		s += string(rune(hex[int(r)/16]))

		if count == 4 {
			z01.PrintRune('\n')
			count = 0
		}
		
		z01.PrintRune(rune(hex[int(r)/16]))
		z01.PrintRune(rune(hex[int(r)%16]))
		count++
		if count != 4  && i != len(arr)-1 {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
	for _, r := range arr {
		if !(r >= 32 && r <= 126) {
			r = '.'
		}
		z01.PrintRune(rune(r))
	}
	z01.PrintRune('\n')
}
