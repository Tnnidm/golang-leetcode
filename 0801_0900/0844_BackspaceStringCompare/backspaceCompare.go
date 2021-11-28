package backspaceCompare

func backspaceCompare(s string, t string) bool {
	s = backspaceProcess(s)
	t = backspaceProcess(t)
	return s == t
}

func backspaceProcess(s string) string {
	stack := []byte{}
	for i := range s {
		if s[i] == '#' {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
