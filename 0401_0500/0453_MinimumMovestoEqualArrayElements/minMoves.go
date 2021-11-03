package minMoves

func minMoves(nums []int) int {
	numMin := nums[0]
	sum := 0
	numsLen := len(nums)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < numMin {
			numMin = nums[i]
		}
	}
	return sum - numMin*numsLen
}
