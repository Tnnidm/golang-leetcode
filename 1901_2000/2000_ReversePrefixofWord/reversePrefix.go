package reversePrefix

func reversePrefix(word string, ch byte) string {
	firstFound := 0
	for i := range word {
		if word[i] == ch {
			firstFound = i
			break
		}
	}
	wordByte := []byte(word)
	for i := 0; i <= firstFound/2; i++ {
		wordByte[i], wordByte[firstFound-i] = wordByte[firstFound-i], wordByte[i]
	}
	return string(wordByte)
}
