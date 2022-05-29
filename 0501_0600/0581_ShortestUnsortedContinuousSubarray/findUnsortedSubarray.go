package findUnsortedSubarray

func findUnsortedSubarray(nums []int) int {
	numsLen := len(nums)
	min, max := nums[numsLen-1], nums[0]
	begin, end := 0, -1
	for i := 0; i < numsLen; i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
		if nums[numsLen-i-1] > min {
			begin = numsLen - i - 1
		} else {
			min = nums[numsLen-i-1]
		}
	}
	return end - begin + 1
}
