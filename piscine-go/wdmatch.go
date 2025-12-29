package piscine

import (
	"fmt"
	"os"
)

func Wdmatch(){
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	newStr := ""


	for _, r := range str1{
		for i := 0; i < len(str2); i++{
			if r == rune(str2[i]) {
				newStr += string(r)
				str2 = str2[i:]
				break
			}

		}
	}
	if newStr == str1 {
		fmt.Println(str1)
		return
	}else {
		return
	}
}