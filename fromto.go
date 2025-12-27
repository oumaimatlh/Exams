package main

//import "fmt"

// import "fmt"

// import "github.com/01-edu/z01"

func Itoaa(n int) string {
	s := ""
	for n > 0 {
		r := n % 10
		s = string(rune(r+'0')) + s
		n /= 10
	}
	return s
}

func FromTo(from int, to int) string {
	if from > 99 || from < 0 {
		return "Invalid\n"
	}
	newStr := ""
	if from < to {
		for i := from; i <= to; i++ {
			if i >= 10 {
				if i == to {
					newStr += Itoaa(i)
					continue
				}
				newStr += Itoaa(i) + "," + " "
			} else {
				if i == to {
					newStr += "0" + Itoaa(i)
					continue
				}
				newStr += "0" + Itoaa(i) + "," + " "
			}
		}
	} else if from > to {
		for i := from; i >= to; i-- {
			if i >= 10 {
				if i == to {
					newStr += Itoaa(i)
					continue
				}
				newStr += Itoaa(i) + "," + " "
			} else {
				if i == to {
					newStr += "0" + Itoaa(i)
					continue
				}
				newStr += "0" + Itoaa(i) + "," + " "
			}
		}
	}else {
		newStr = Itoaa(from)
	}

	return newStr + "\n"
}
