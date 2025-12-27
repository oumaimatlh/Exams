package main

func LastWord(s string) string {
	// Words
	words := []string{}
	word := ""
	check := false
	for _, r := range s {
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

	if len(words) ==0 {
		return  "\n"
	}
	return words[len(words)-1]+"\n"
}
