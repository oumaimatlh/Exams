package piscine

import (
	///"fmt"
	"os"

	"github.com/01-edu/z01"
)

func LowerCase(r rune) rune {
	r = r + 32
	return r
}

func UpperCase(r rune) rune {
	r = r - 32
	return r
}

func main() {
	for j := 0; j < len(os.Args); j++ {
		s := os.Args[j]
		str := ""

		for i := 0; i < len(s); i++ {
			if i == len(s)-1 && s[i] >= 'a' && s[i] <= 'z' {
				str += string(UpperCase(rune(s[i])))
				break
			}

			if i < len(s)-1 && s[i+1] == ' ' {
				if s[i] >= 'a' && s[i] <= 'z' {
					str += string(UpperCase(rune(s[i])))
				} else {
					str += string(s[i])
				}
			} else {
				if s[i] >= 'A' && s[i] <= 'Z' {
					str += string(LowerCase(rune(s[i])))
				} else {
					str += string(s[i])
				}
			}
		}

		for _, r := range str {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
