package countCharacters

func countCharacters(words []string, chars string) int {
	charMap := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		charMap[chars[i]-'a']++
	}
	ans := 0
	charMapCopy := make([]int, 26)
	for _, word := range words {
		for j := 0; j < 26; j++ {
			charMapCopy[j] = charMap[j]
		}
		flag := true
		for _, v := range word {
			charMapCopy[v-'a']--
			if charMapCopy[v-'a'] < 0 {
				flag = false
				break
			}
		}
		if flag {
			ans = ans + len(word)
		}
	}
	return ans
}
