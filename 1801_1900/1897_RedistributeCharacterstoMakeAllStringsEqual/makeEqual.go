package makeEqual

func makeEqual(words []string) bool {
	wordNum := len(words)
	if wordNum == 1 {
		return true
	}
	charMap := make([]int, 26)
	for i := range words {
		for _, v := range words[i] {
			charMap[v-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if charMap[i]%wordNum != 0 {
			return false
		}
	}
	return true
}
