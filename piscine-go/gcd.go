package piscine

import "fmt"

func Gcd(a, b uint) uint {
	diva := []uint{}
	divb := []uint{}

	for i := 1; i <= int(a); i++ {
		if a%uint(i) == 0 {
			diva = append(diva, uint(i))
		}
	}

	for i := 1; i <= int(b); i++ {
		if b%uint(i) == 0 {
			divb = append(divb, uint(i))
		}
	}

	// GCD
	
	gcd := -1
	for i := 0; i < len(diva); i++ {
		for j := 0; j < len(divb); j++ {
			if diva[i] == divb[j] && int(diva[i]) > gcd {
				gcd = int(diva[i])
			}
		}
	}

	fmt.Println(diva)
	fmt.Println(divb)
	return uint(gcd)
}
