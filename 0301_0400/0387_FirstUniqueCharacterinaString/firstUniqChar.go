package firstUniqChar

func firstUniqChar(s string) int {
	charMap := make([]int, 26)
	for i := 0; i < len(s); i++ {
		charMap[s[i]-'a'] += 1
	}
	for i := 0; i < len(s); i++ {
		if charMap[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
