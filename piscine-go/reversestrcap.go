package piscine 
import "os"
func ReversestrCap(){
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	words := []string{}
	word := ""
    check := false 
	for _, r := range s {
        if r != ' ' {
           word += string(r)
		   check = true
		}else if check  {
			words = append(words, word)
			word = ""
			check = false
		}
	}

	if len(word) != 0 {
		words = append(words, word)
	}

/* 	for _, r := range words {
		
	} */
}