package removeOuterParentheses

import "strings"

func removeOuterParentheses(s string) string {
	leftRight := 0
	result := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftRight++
			if leftRight != 1 {
				result.WriteByte(s[i])
			}
		} else {
			leftRight--
			if leftRight != 0 {
				result.WriteByte(s[i])
			}
		}
	}
	return result.String()
}
