package pivotIndex

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	preSum := 0
	for i, num := range nums {
		if preSum<<1+num == sum {
			return i
		}
		preSum += num
	}
	return -1
}
