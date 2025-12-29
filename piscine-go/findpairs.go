package piscine

import (
	"fmt"
	"os"
	"strconv"
)

func findPairs() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	sum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	s := os.Args[1]
	n := []int{}

	if !(s[0] == '[' && s[len(s)-1] == ']') {
		fmt.Println("Invalid input.")
		return
	}

	newS := s[1 : len(s)-1]
	p := ""
	for i := 0; i < len(newS); i++ {
		if newS[i] != ',' {
			p += string(newS[i])
		} else {
			i++
			number, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Invalid number: ", p)
				return
			}
			n = append(n, number)
			p = ""

		}
	}
	if len(p) != 0 {
		number, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Invalid number: ", p)
			return
		}
		n = append(n, number)
	}

	//---------------------------------------------//
	output := "Pairs with sum " + os.Args[2] + ": ["
	check := false

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == sum {
				output += "[" + strconv.Itoa(i) + " " + strconv.Itoa(j) + "] "
				check = true
			}
		}
	}
	if !check {
		fmt.Println("No pairs found.")

		return
	}
	output = output[:len(output)-1]
	fmt.Println(output + "]")
}
