package sortedSquares

func sortedSquares(nums []int) []int {
	numsLen := len(nums)
	i, j := 0, numsLen-1
	result := make([]int, numsLen)
	power1 := nums[i] * nums[i]
	power2 := nums[j] * nums[j]
	for index := numsLen - 1; index >= 0; index-- {
		power1 = nums[i] * nums[i]
		power2 = nums[j] * nums[j]
		if power1 >= power2 {
			result[index] = power1
			i++
		} else {
			result[index] = power2
			j--
		}
	}
	return result
}
