package decrypt

func decrypt(code []int, k int) []int {
	codeNum := len(code)
	result := make([]int, codeNum)
	if k == 0 {
		return result
	} else if k > 0 {
		for i := 0; i < codeNum; i++ {
			for j := 1; j <= k; j++ {
				result[i] = result[i] + code[(i+j+codeNum)%codeNum]
			}
		}
	} else {
		for i := 0; i < codeNum; i++ {
			for j := -1; j >= k; j-- {
				result[i] = result[i] + code[(i+codeNum+j)%codeNum]
			}
		}
	}
	return result
}
