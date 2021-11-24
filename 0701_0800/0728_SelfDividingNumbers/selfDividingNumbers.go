package selfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for num := left; num <= right; num++ {
		selfDividingFlag := true
		for n := num; n != 0; n = n / 10 {
			if n%10 == 0 || num%(n%10) != 0 {
				selfDividingFlag = false
				break
			}
		}
		if selfDividingFlag {
			result = append(result, num)
		}
	}
	return result
}
