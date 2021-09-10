package fib

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了71.26%的用户
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		fib1 := 0
		fib2 := 1
		fib3 := 0
		for i := 2; i <= n; i++ {
			fib3 = fib1 + fib2
			fib1 = fib2
			fib2 = fib3
		}
		return fib3
	}
}
