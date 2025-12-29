package piscine

func Slice(a []string, nbrs ...int) []string {
	newSlice := []string{}

	for i := 0; i < len(nbrs); i++ {
		index := nbrs[i]

		if index < 0 {
			index = -1 * index
			if index > len(a) || (i+1 < len(nbrs) && nbrs[i+1] > len(a)) {
				return []string{}
			}
			if len(nbrs) == 1 {
				newSlice = append(newSlice, a[len(a)-index:]...)
			} else {
				if i+1 < len(nbrs) {
					if -index > nbrs[i+1] {
						return []string(nil)
					} else {
						newSlice = append(newSlice, a[len(a)-index:len(a)-(-1*nbrs[i+1])]...)
					}
				}
			}

		} else {

			if index > len(a) || (i+1 < len(nbrs) && nbrs[i+1] > len(a)) {
				return []string{}
			}
			if index > len(a) || (i+1 < len(nbrs) && nbrs[i+1] > len(a)) {
				return []string{}
			}
			if len(nbrs) == 1 {
				newSlice = append(newSlice, a[index:]...)
			} else {
				if i+1 < len(nbrs) {
					if index > nbrs[i+1] {
						return []string(nil)
					} else {
						newSlice = append(newSlice, a[index:nbrs[i+1]]...)
					}
				}
			}
		}

	}

	return newSlice
}
