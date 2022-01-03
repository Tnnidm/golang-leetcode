package checkString

func checkString(s string) bool {
	sLen := len(s)
	for i := 1; i < sLen; i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
