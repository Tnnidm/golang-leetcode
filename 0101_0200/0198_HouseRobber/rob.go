package rob

func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	if numsLen == 2 {
		return max(nums[0], nums[1])
	}
	dp1, dp2 := 0, 0
	for _, num := range nums {
		dp1, dp2 = dp2, max(dp1+num, dp2)
	}
	return max(dp1, dp2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
