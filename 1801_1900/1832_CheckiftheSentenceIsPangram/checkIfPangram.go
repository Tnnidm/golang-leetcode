package checkIfPangram

func checkIfPangram(sentence string) bool {
	charMap := make([]int, 26)
	for _, char := range sentence {
		charMap[char-'a']++
	}
	for i := 0; i < 26; i++ {
		if charMap[i] == 0 {
			return false
		}
	}
	return true
}
