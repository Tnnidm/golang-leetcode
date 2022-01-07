package totalMoney

func totalMoney(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result = result + i%7 + 1 + i/7
	}
	return result
}
