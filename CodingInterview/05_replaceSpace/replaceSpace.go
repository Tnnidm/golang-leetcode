package replaceSpace

// // 直接调用strings.ReplaceAll
// func replaceSpace(s string) string {
// 	return strings.ReplaceAll(s, " ", "%20")
// }

// // 调用strings.Builder
// func replaceSpace(s string) string {
// 	sLen := len(s)
// 	if sLen == 0 {
// 		return s
// 	}
// 	ans := strings.Builder{}
// 	for i := 0; i < sLen; i++ {
// 		if s[i] == ' ' {
// 			ans.WriteString("%20")
// 		} else {
// 			ans.WriteByte(s[i])
// 		}
// 	}
// 	return ans.String()
// }

// 手动实现
// 时间复杂度O(n), 空间复杂度O(n)
func replaceSpace(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return s
	}
	newSLen := sLen
	for i := 0; i < sLen; i++ {
		if s[i] == ' ' {
			newSLen += 2
		}
	}
	if newSLen == sLen {
		return s
	}
	newS := make([]byte, newSLen)
	j := 0
	for i := 0; i < sLen; i++ {
		if s[i] == ' ' {
			newS[j] = '%'
			j++
			newS[j] = '2'
			j++
			newS[j] = '0'
			j++
		} else {
			newS[j] = s[i]
			j++
		}
	}
	return string(newS)
}
