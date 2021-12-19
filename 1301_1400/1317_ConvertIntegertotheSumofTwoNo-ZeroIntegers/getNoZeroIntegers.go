package getNoZeroIntegers

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		if noZero(i) && noZero(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}

func noZero(num int) bool {
	for num != 0 {
		if num%10 == 0 {
			return false
		}
		num = num / 10
	}
	return true
}
