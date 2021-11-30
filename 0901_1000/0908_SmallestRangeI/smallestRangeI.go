package smallestRangeI

func smallestRangeI(nums []int, k int) int {
	maxValue, minValue := nums[0], nums[0]
	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
	}
	if maxValue-minValue <= 2*k {
		return 0
	} else {
		return maxValue - minValue - 2*k
	}
}
