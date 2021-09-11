package getMaximumGenerated

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了92.63%的用户
func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	} else {
		numsMax := 1
		nums := make([]int, n+1)
		nums[0] = 0
		nums[1] = 1
		for i := 2; i <= n; i++ {
			if i%2 == 0 {
				nums[i] = nums[i/2]
			} else {
				nums[i] = nums[i/2] + nums[i/2+1]
			}
			if nums[i] > numsMax {
				numsMax = nums[i]
			}
		}
		return numsMax
	}
}
