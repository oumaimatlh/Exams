package main

func IsPrime(n int) bool {
	count := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	if count == 2 {
		return true
	}
	return false
}

func FindPrevPrime(nb int) int {
	PrevPrime := 0

	for i := nb; i >= 2 ; i-- {
		if IsPrime(i) {
			PrevPrime = i
			break
		}
	}
	if PrevPrime == 0 {
		return  0
	}
	return PrevPrime
}
