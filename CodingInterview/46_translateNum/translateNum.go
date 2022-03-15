package translateNum

// // 第一种方法是转换成字符串，然后靠递归计算得到结果
// func translateNum(num int) int {
// 	numStr := strconv.Itoa(num)
// 	return translate(numStr)
// }

// func translate(numStr string) int {
// 	numStrLen := len(numStr)
// 	if numStrLen <= 1 {
// 		return 1
// 	} else if numStrLen > 1 && (numStr[0] == '1' || (numStr[0] == '2' && numStr[1] <= '5')) {
// 		return translate(numStr[1:]) + translate(numStr[2:])
// 	} else {
// 		return translate(numStr[1:])
// 	}
// }

// 第二种方法是直接按位遍历这个数，本质上类似跳楼梯问题，是个变种斐波那契数列
func translateNum(num int) int {
	var thisDigit, lastDigit int
	var twoDigitNum int
	lastDigit = num % 10
	thisNum := 1
	lastNum := 1
	for num != 0 {
		num /= 10
		thisDigit = num % 10
		twoDigitNum = thisDigit*10 + lastDigit
		if twoDigitNum >= 10 && twoDigitNum <= 25 {
			thisNum, lastNum = thisNum+lastNum, thisNum
		} else {
			lastNum = thisNum
		}
		lastDigit = thisDigit
	}
	return thisNum
}
