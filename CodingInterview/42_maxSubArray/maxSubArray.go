package maxSubArray

func maxSubArray(nums []int) int {
	count, maxCount := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		count = max(count+nums[i], nums[i])
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
