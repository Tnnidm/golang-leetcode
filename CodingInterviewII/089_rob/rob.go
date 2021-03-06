package rob

func rob(nums []int) int {
	numsLen := len(nums)
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
