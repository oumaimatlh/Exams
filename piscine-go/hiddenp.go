package piscine

import (
	//"fmt"
//	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func HiddenP() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	count := 0

	j:= 0
	for i := 0; i < len(s1); i++ {
		for j := j; j < len(s2); j++ {
			if s1[i] == s2[j] {
				count++
				j = j + 1
				break
			}
		}
	}
	if count == len(s1) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
