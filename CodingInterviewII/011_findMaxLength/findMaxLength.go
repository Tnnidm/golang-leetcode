package findMaxLength

func findMaxLength(nums []int) int {
	firstIndex := map[int]int{}
	firstIndex[0] = -1
	maxLen := 0
	preSum := 0
	for i, num := range nums {
		if num == 0 {
			preSum--
		} else {
			preSum++
		}
		if _, ok := firstIndex[preSum]; ok {
			maxLen = max(maxLen, i-firstIndex[preSum])
		} else {
			firstIndex[preSum] = i
		}
	}
	return maxLen
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
