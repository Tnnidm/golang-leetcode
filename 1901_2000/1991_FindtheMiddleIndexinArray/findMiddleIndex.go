package findMiddleIndex

func findMiddleIndex(nums []int) int {
	numsLen := len(nums)
	totalSum := 0
	for i := 0; i < numsLen; i++ {
		totalSum += nums[i]
	}
	leftSum := 0
	for middleIndex := 0; middleIndex < numsLen; middleIndex++ {
		if 2*leftSum+nums[middleIndex] == totalSum {
			return middleIndex
		}
		leftSum += nums[middleIndex]
	}
	return -1
}
