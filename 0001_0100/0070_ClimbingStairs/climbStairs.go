package climbStairs

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了99.97%的用户
func climbStairs(n int) int {
	var result int
	if n == 1 {
		result = 1
	} else if n == 2 {
		result = 2
	} else {
		methods_n_minus_2 := 1
		methods_n_minus_1 := 2
		for i := 3; i <= n; i++ {
			result = methods_n_minus_1 + methods_n_minus_2
			methods_n_minus_2 = methods_n_minus_1
			methods_n_minus_1 = result
		}
	}
	return result
}
