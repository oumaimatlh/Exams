package piscine

import (
	"fmt"
	"os"
	//"strings"
)

func RevWstr() {
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
		}
	}
	if len(word) != 0 {
		words = append(words, word)
	}

	s := ""
	for i := len(words)-1; i >= 0; i-- {
		s += words[i] + " "
	}
	fmt.Println(s)
}
