package longestNiceSubarray

func longestNiceSubarray(nums []int) int {
	var maxLen, orResult, sumResult int
	left := 0
	for right := range nums {
		for orResult|nums[right] != sumResult+nums[right] {
			orResult -= nums[left]
			sumResult -= nums[left]
			left++
		}
		orResult = orResult | nums[right]
		sumResult += nums[right]
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
