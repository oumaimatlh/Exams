package main



func RepeatAlpha(s string) string {
	newStr := ""
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z') {
			newStr += string(r)
			continue
		}

		repeat := 0
		if r >= 'a' && r <= 'z' {
			repeat = int(r) - 96

			for i := 1; i <= repeat; i++ {
				newStr += string(r)
			}
		}

		if r >= 'A' && r <= 'Z' {
			repeat = int(r) - 64
			for i := 1; i <= repeat; i++ {
				newStr += string(r) 
			}
		}
	}
	return newStr
}
