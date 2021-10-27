package isAnagram

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	} else {
		table := [26]int{}
		for i := 0; i < sLen; i++ {
			table[int(s[i]-'a')]++
			table[int(t[i]-'a')]--
		}
		for i := 0; i < 26; i++ {
			if table[i] != 0 {
				return false
			}
		}
		return true
	}
}
