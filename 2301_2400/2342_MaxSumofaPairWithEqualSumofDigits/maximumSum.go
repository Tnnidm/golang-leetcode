package maximumSum

func maximumSum(nums []int) int {
	result := -1
	digitSum2MaxNum := map[int]int{}
	for _, num := range nums {
		digitsum := digitSum(num)
		if digitSum2MaxNum[digitsum] > 0 {
			result = max(result, digitSum2MaxNum[digitsum]+num)
			digitSum2MaxNum[digitsum] = max(digitSum2MaxNum[digitsum], num)
		} else {
			digitSum2MaxNum[digitsum] = num
		}
	}
	return result
}

func digitSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
