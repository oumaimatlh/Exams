package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	sliceAlternate := []int{}

	for i := 0; i < len(slice1)+len(slice2); i++ {
		if len(slice1) > len(slice2) {
			if i < len(slice1) {
				sliceAlternate = append(sliceAlternate, slice1[i])
			}
			if i < len(slice2) {
				sliceAlternate = append(sliceAlternate, slice2[i])
			}
			//

		} else {
			if i < len(slice2) {
				sliceAlternate = append(sliceAlternate, slice2[i])
			}
			if i < len(slice1) {
				sliceAlternate = append(sliceAlternate, slice1[i])
			}
		}
	}
	return sliceAlternate
}
