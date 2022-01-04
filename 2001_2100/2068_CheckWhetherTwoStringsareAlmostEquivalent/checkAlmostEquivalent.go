package checkAlmostEquivalent

func checkAlmostEquivalent(word1 string, word2 string) bool {
	charList := make([]int, 26)
	for _, v := range word1 {
		charList[v-'a']++
	}
	for _, v := range word2 {
		charList[v-'a']--
	}
	for _, v := range charList {
		if v > 3 || v < -3 {
			return false
		}
	}
	return true
}
