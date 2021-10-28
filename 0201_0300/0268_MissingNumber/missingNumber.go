package missingNumber

func missingNumber(nums []int) int {
	numsLen := len(nums)
	sum := (numsLen + 1) * numsLen / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
