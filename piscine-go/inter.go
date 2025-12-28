package piscine

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func check(r byte, s string) bool {
	for _, k := range s {
		if k == rune(r) {
			return true
		}
	}
	return false
}

func Inter() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	s := ""

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				if !check(s1[i], s) {
					s += string(s1[i])
				}
				j++
				break
			}
		}
	}


	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
