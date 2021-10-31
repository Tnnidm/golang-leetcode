package findTheDifference

// // 字典法
// func findTheDifference(s string, t string) byte {
// 	charMap := make([]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		charMap[s[i]-'a'] += 1
// 	}
// 	var result byte
// 	for i := 0; i < len(t); i++ {
// 		if charMap[t[i]-'a'] <= 0 {
// 			result = t[i]
// 		} else {
// 			charMap[t[i]-'a'] -= 1
// 		}
// 	}
// 	return result
// }

// 其实可以直接全部加起来，后一个减去前一个的值就是多的那一个
func findTheDifference(s string, t string) byte {
	var sum int
	for i := 0; i < len(t); i++ {
		sum += int(t[i])
	}
	for i := 0; i < len(s); i++ {
		sum -= int(s[i])
	}
	return byte(sum)
}
