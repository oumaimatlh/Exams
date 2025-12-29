package piscine


func SaveAndMiss(arg string, num int) string {
	if num <=0 {
		return arg
	}
	newStr := ""
	count := 0

	for i := 0; i < len(arg); i++{

		if count != num {
			count++
			newStr += string(arg[i])

		} else {
			i += num-1
			count = 0
		}
	}
	return newStr
}
