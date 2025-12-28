package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	slices := [][]int{}
	i := 0


	for {
		if i+size < len(slice)-1 {

			slices = append(slices, slice[i:i+size])
			i += size

		} else {
			slices = append(slices, slice[i:])
			break
		}
	}
	fmt.Println(slices)
}
