package maxAscendingSum

func maxAscendingSum(nums []int) int {
	maxSubSum, subSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			subSum += nums[i]
		} else {
			subSum = nums[i]
		}
		if subSum > maxSubSum {
			maxSubSum = subSum
		}
	}
	return maxSubSum
}
