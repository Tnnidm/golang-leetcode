package findAnagrams

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return []int{}
	}
	charMap := make([]int, 26)
	for i := 0; i < pLen; i++ {
		charMap[p[i]-'a']++
		charMap[s[i]-'a']--
	}
	result := []int{}
	if isAllZero(charMap) {
		result = append(result, 0)
	}
	for i := pLen; i < sLen; i++ {
		charMap[s[i]-'a']--
		charMap[s[i-pLen]-'a']++
		if isAllZero(charMap) {
			result = append(result, i-pLen+1)
		}
	}
	return result
}

func isAllZero(arr []int) bool {
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}
