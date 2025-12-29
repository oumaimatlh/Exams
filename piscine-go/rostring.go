package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := []string{}
	word := ""
	check := false
	str := os.Args[1]
	for _, r := range str {
		if r != ' ' { 
			word += string(r)
			check = true
		} else if check {
			words = append(words, word)
			word = ""

			check = false
		}
	}
	if len(word) != 0 {
		words = append(words, word)
	}
	fmt.Println(strings.Join(words, "||"))
	

	s := ""
	
	// for i := 0 ; i < len(words[1:]); i++ {
	// 	s += words[i] + " "
	// }
	
	// s += words[0]
	fmt.Println(s)
	

}
