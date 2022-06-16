package multiply

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Len := len(num1)
	num2Len := len(num2)
	resultLen := num1Len + num2Len
	resultBit := make([]int, resultLen)
	for k := 0; k < resultLen; k++ {
		for i := max(0, k+1-num2Len); i <= k && i < num1Len; i++ {
			resultBit[resultLen-1-k] += int(num1[num1Len-1-i]-'0') * int(num2[num2Len-1-(k-i)]-'0')
		}
	}
	for k := resultLen - 1; k >= 1; k-- {
		if resultBit[k] >= 10 {
			resultBit[k-1] += resultBit[k] / 10
			resultBit[k] = resultBit[k] % 10
		}
	}
	if resultBit[0] == 0 {
		resultBit = resultBit[1:]
		resultLen--
	}
	result := make([]byte, resultLen)
	for index, value := range resultBit {
		result[index] = '0' + uint8(value)
	}
	return string(result)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
