package arrangeCoins

func arrangeCoins(n int) int {
	var i int
	for i = 1; i <= n; i++ {
		n -= i
	}
	return i - 1
}
