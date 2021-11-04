package findMaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	begin, maxCount := 0, 0
	for i := 0; i < len(nums); i++ {
		if (i == 0 || nums[i-1] == 0) && nums[i] == 1 {
			begin = i
		}
		if (i == len(nums)-1 || nums[i+1] == 0) && nums[i] == 1 {
			maxCount = max(maxCount, i-begin+1)
		}
	}
	return maxCount
}

func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
