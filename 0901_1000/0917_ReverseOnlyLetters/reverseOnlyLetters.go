package reverseOnlyLetters

func reverseOnlyLetters(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i < j; i, j = i+1, j-1 {
		for !isLetter(sByte[i]) && i < j {
			i++
		}
		for !isLetter(sByte[j]) && i < j {
			j--
		}
		sByte[i], sByte[j] = sByte[j], sByte[i]
	}
	return string(sByte)
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
