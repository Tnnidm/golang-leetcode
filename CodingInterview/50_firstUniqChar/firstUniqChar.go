package firstUniqChar

func firstUniqChar(s string) byte {
	var countMap [26]int
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		countMap[s[i]-'a']++
	}
	for i := 0; i < sLen; i++ {
		if countMap[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
