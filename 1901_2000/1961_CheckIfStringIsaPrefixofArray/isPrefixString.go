package isPrefixString

func isPrefixString(s string, words []string) bool {
	start := 0
	sLen := len(s)
	for i := 0; i < len(words) && start < sLen; i++ {
		if start+len(words[i]) <= sLen && s[start:start+len(words[i])] != words[i] {
			return false
		}
		start += len(words[i])
	}
	return start == sLen
}
