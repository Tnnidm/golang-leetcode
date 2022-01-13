package makeFancyString

import "strings"

func makeFancyString(s string) string {
	result := strings.Builder{}
	continueNum := 0
	result.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			result.WriteByte(s[i])
			continueNum = 0
		} else if continueNum == 0 {
			result.WriteByte(s[i])
			continueNum++
		}
	}
	return result.String()
}
