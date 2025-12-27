package main

// import  "github.com/01-edu/z01"
func HashCode(dec string) string {
	/*
		(ASCII of current character + size of the string) % 127,
		ensuring the result falls within the ASCII range of 0 to 127.

	*/
	hashCode := ""
	for _, r := range dec {

		hash := (int(r) + len(dec)) % 127
		if hash >= 0 && hash <= 31 || hash == 127 {
			hash = hash + 33
		}
		hashCode += string(rune(hash))

	}
	return hashCode
}
