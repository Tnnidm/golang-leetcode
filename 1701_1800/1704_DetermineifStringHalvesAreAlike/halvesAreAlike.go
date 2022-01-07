package halvesAreAlike

func halvesAreAlike(s string) bool {
	vowelNum := 0
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelNum++
		}
	}
	for i := sLen / 2; i < sLen; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelNum--
		}
	}
	return vowelNum == 0
}
