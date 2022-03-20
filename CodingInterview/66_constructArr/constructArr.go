package constructArr

func constructArr(a []int) []int {
	aLen := len(a)
	result := make([]int, aLen)
	if aLen == 0 {
		return result
	}
	result[0] = 1
	for i := 1; i < aLen; i++ {
		result[i] = result[i-1] * a[i-1]
	}
	assist := 1
	for i := aLen - 2; i >= 0; i-- {
		assist *= a[i+1]
		result[i] *= assist
	}
	return result
}
