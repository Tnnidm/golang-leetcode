package myAtoi

func myAtoi(s string) int {
	byteArr := []byte(s)
	// 处理开头的空格
	index := 0
	for index < len(byteArr) && byteArr[index] == ' ' {
		index++
	}
	byteArr = byteArr[index:]
	negFlagFlag := false
	if len(byteArr) > 0 {
		if byteArr[0] == '-' {
			negFlagFlag = true
			byteArr = byteArr[1:]
		} else if byteArr[0] == '+' {
			byteArr = byteArr[1:]
		}
	}
	result := 0
	for index = 0; index < len(byteArr) && byteArr[index] <= '9' && byteArr[index] >= '0'; index++ {
		result = 10*result + int(byteArr[index]-'0')
		if result > 2147483647 {
			if negFlagFlag {
				return -2147483648
			} else {
				return 2147483647
			}
		}
	}
	if negFlagFlag {
		result = -result
	}
	return result
}
