package isAlienSorted

func isAlienSorted(words []string, order string) bool {
	charMap := make([]int, 26)
	for i := 0; i < 26; i++ {
		charMap[order[i]-'a'] = i
	}
	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i], charMap) {
			return false
		}
	}
	return true
}

func less(x, y string, charMap []int) bool {
	xLen := len(x)
	yLen := len(y)
	for i := 0; i < xLen && i < yLen; i++ {
		if charMap[x[i]-'a'] > charMap[y[i]-'a'] {
			return false
		} else if charMap[x[i]-'a'] < charMap[y[i]-'a'] {
			return true
		}
	}
	return xLen <= yLen
}
