package maxProduct

func maxProduct(nums []int) int {
	result := 0
	maxValue, minValue, result := nums[0], nums[0], nums[0]
	var product1, product2 int
	for i := 1; i < len(nums); i++ {
		product1 = maxValue * nums[i]
		product2 = minValue * nums[i]
		maxValue = max(max(product1, product2), nums[i])
		minValue = min(min(product1, product2), nums[i])
		result = max(result, maxValue)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
