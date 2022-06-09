package decodeString

import "strings"

func decodeString(s string) string {
	result := strings.Builder{}
	sLen := len(s)
	leftIndex := -1
	leftNum := 0
	number := 0
	for i := 0; i < sLen; i++ {
		if leftNum == 0 {
			if s[i] >= '0' && s[i] <= '9' {
				number = number*10 + int(s[i]-'0')
				continue
			} else if s[i] >= 'a' && s[i] <= 'z' {
				result.WriteByte(s[i])
				continue
			}
		}
		if s[i] == '[' {
			if leftIndex == -1 {
				leftIndex = i
			}
			leftNum++
			continue
		}
		if s[i] == ']' {
			leftNum--
			if leftNum == 0 {
				temp := decodeString(s[leftIndex+1 : i])
				for j := 0; j < number; j++ {
					result.WriteString(temp)
				}
				leftIndex = -1
				number = 0
			}
			continue
		}
	}
	return result.String()
}
