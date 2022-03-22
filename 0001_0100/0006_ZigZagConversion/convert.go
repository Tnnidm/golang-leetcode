package convert

func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows == 1 {
		return s
	}
	groupNum := 2*numRows - 2
	result := make([]byte, sLen)
	index := 0
	for i := 0; i < sLen; i += groupNum {
		result[index] = s[i]
		index++
	}
	for row := 1; row < groupNum/2; row++ {
		for i := row; i < sLen; i += groupNum {
			result[index] = s[i]
			index++
			if i+groupNum-2*row < sLen {
				result[index] = s[i+groupNum-2*row]
				index++
			}
		}
	}
	for i := groupNum / 2; i < sLen; i += groupNum {
		result[index] = s[i]
		index++
	}
	return string(result)
}
