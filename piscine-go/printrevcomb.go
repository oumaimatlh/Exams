package piscine

import (
	//"fmt"

	"github.com/01-edu/z01"
)

func PrintRevComb() {
	/*
		999 or 000 are not valid combinations because the digits are not different.
		789 should not be shown because the first digit is not greater than the second.

	*/

	for i := '9'; i >= '2'; i-- {
		for j := '8'; j >= '1'; j-- {
			for q := '7'; q >= '0'; q-- {
				if i > j && j > q {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(q)

					if !(i == '2' && j == '1' && q == '0'){
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
			}
		}
	}
}
