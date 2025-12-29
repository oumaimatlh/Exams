package piscine

import (
	"fmt"
	"os"
)

func checkRune(r rune, s string) bool {
	for _, k := range s {
		if k == r {
			return true
		}
	}
	return false
}

func Union() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	newStr := ""


	for _, r := range str1 {
		if !checkRune(r, newStr) {
			newStr += string(r)
		}
	}
	for _, r := range str2 {
		if !checkRune(r, newStr) {
			newStr += string(r)
		}
	}

	fmt.Println(newStr)
}
