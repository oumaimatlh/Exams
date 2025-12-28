package piscine

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 &&  len(str2) == 0 {
		return  -1
	}
	count := 0
	for _, r := range str1 {
		check := false

		for _, k := range str2 {
			if r == k {
				check = true
				break
			}
		}
		if !check {
			count++
		}
	}
	for _, r := range str2 {
		check := false

		for _, k := range str1 {
			if r == k {
				check = true
				break
			}
		}
		if !check {
			count++
		}
	}
	return count
}
