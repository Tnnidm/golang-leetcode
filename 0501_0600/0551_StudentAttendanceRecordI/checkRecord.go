package checkRecord

func checkRecord(s string) bool {
	LCount, ACount := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			if i == 0 || s[i-1] != 'L' {
				LCount = 1
			} else {
				LCount += 1
			}
			if LCount == 3 {
				return false
			}
		} else if s[i] == 'A' {
			ACount += 1
			if ACount == 2 {
				return false
			}
		}
	}
	return true
}

// // 可以直接使用标准库一行秒杀
// func checkRecord(s string) bool {
// 	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
// }
