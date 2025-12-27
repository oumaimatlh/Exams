package main

import "os"

func ExpandStr() string {
	if len(os.Args) != 2 {
		return ""
	}
	str := os.Args[1]
	words := []string{}
	word := ""
	check := false

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

	newStr := ""
	for i, r := range words {
		if i == len(words)-1 {
			newStr += r
			continue

		}
		newStr += r + "   "
	}
	return newStr
}
