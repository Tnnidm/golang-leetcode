package areOccurrencesEqual

func areOccurrencesEqual(s string) bool {
	charMap := make([]int, 26)
	for i := range s {
		charMap[s[i]-'a']++
	}
	showTime := charMap[s[0]-'a']
	for i := 0; i < 26; i++ {
		if charMap[i] != showTime && charMap[i] != 0 {
			return false
		}
	}
	return true
}
