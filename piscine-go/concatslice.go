package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 && len(slice2) == 0{
		return  []int(nil)
	}
	ConcatSLice := []int{}

	for i := 0; i < len(slice1); i++ {
		ConcatSLice = append(ConcatSLice, slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		ConcatSLice = append(ConcatSLice, slice2[i])
	}

	return ConcatSLice
}
