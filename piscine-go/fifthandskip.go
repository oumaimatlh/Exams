package piscine

func FifthAndSkip(str string) string {
	if len(str) == 0  {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input"
	}
	s:= ""
	
	for _, r := range str {
		if r != ' '{
			s += string(r)
		}
	}
	newStr := ""
	count := 0

	for i, r := range s {
		if count != 5 {
			newStr += string(r)
			count++

		} else {
			newStr += " "
			count = 0
			i++
		}
	}
	return newStr + "\n"
}
