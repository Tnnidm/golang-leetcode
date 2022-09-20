package removeDuplicateLetters

func removeDuplicateLetters(s string) string {
	left := make([]int, 26)
	for i := range s {
		left[s[i]-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		if !inStack[s[i]-'a'] {
			stackLen := len(stack)
			for stackLen > 0 && s[i] < stack[stackLen-1] && left[stack[stackLen-1]-'a'] > 0 {
				stackLen--
				inStack[stack[stackLen]-'a'] = false
				stack = stack[:stackLen]
			}
			stack = append(stack, s[i])
			inStack[s[i]-'a'] = true
		}
		left[s[i]-'a']--
	}
	return string(stack)
}
