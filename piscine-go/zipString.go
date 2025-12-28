package piscine

func ZipString(s string) string {
	newStr := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			count := 1
			for i < len(s)-1 && s[i] == s[i+1] {
				count++
				i++
			}

			newStr += Itoa(count) + string(s[i])
		} else {
			newStr += string(s[i])
		}
	}

	return newStr
}
