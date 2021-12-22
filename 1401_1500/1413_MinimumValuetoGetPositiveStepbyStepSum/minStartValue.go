package minStartValue

func minStartValue(nums []int) int {
	sum := 0
	minSum := nums[0]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if sum < minSum {
			minSum = sum
		}
	}
	if minSum >= 1 {
		return 1
	} else {
		return 1 - minSum
	}
}
