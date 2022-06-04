package productExceptSelf

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	result[0] = 1
	for i := 1; i < numsLen; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	temp := 1
	for i := numsLen - 2; i >= 0; i-- {
		temp *= nums[i+1]
		result[i] *= temp
	}
	return result
}
