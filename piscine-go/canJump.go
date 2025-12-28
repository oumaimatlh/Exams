package piscine

func CanJump(n []uint) bool {
	if len(n) == 0 {
		return false
	}
	if len(n) == 1 {
		return true
	}

	pos := 0
	last := len(n) - 1

	for {
		jump := int(n[pos])

		if jump == 0 {
			return false
		}

		pos += jump

		if pos == last {
			return true
		}
		if pos > last {
			return false
		}
	}
}
