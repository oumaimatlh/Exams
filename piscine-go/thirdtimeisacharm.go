package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return  "\n"
	}
	thirdCharm := ""
	count := 0
	for _, r := range str {
		if count == 2 {
			count = 0
			thirdCharm += string(r)
		}else {
			count++
		}
	}
	return  thirdCharm + "\n"
}