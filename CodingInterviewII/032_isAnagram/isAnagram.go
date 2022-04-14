package isAnagram

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if s == t || sLen != tLen {
		return false
	}
	statistic := make([]int, 26)
	for i := 0; i < sLen; i++ {
		statistic[s[i]-'a']++
	}
	for i := 0; i < tLen; i++ {
		statistic[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if statistic[i] != 0 {
			return false
		}
	}
	return true
}
