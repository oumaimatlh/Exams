package main


func CanJump(n []uint) bool {
	
	valid := false
	
	for i := 0; i < len(n); i++ {
		pos := i
		jump := n[i]

		if  i != len(n)-1 {
			pos = int(jump)
			i += pos+1
		}else {
			valid = true
			//break
		}
	}
	
	return  valid
}