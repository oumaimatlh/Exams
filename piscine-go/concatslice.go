package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	ConcatSLice := []int{}

	for i := 0; i < len(slice1); i++ {
		ConcatSLice = append(ConcatSLice, slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		ConcatSLice = append(ConcatSLice, slice2[i])
	}

	return ConcatSLice
}
