package kLengthApart

func kLengthApart(nums []int, k int) bool {
	numsLen := len(nums)
	lastOne := -1
	for i := 0; i < numsLen; i++ {
		if nums[i] == 1 {
			if lastOne != -1 && i-lastOne-1 < k {
				return false
			}
			lastOne = i
		}
	}
	return true
}
