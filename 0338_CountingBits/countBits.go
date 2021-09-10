package countBits

// 执行用时：4 ms, 在所有 Go 提交中击败了89.10%的用户
// 内存消耗：4.6 MB, 在所有 Go 提交中击败了70.03%的用户
func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	} else {
		result := make([]int, n+1)
		result[0] = 0
		result[1] = 1
		counter := 2
		for i := 2; i <= n; i++ {
			if i/counter == 2 {
				counter = i
			}
			result[i] = result[i-counter] + 1
		}
		return result
	}
}
