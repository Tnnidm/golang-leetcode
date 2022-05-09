package rob

func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	return max(robCore(nums[1:], numsLen-1), robCore(nums[:numsLen-1], numsLen-1))
}

func robCore(nums []int, numsLen int) int {
	robNum := [2]int{nums[0], 0}
	for i := 1; i < numsLen; i++ {
		robNum[0], robNum[1] = robNum[1]+nums[i], max(robNum[0], robNum[1])
	}
	return max(robNum[1], robNum[0])
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
