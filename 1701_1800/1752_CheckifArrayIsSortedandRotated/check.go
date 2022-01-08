package check

func check(nums []int) bool {
	numsLen := len(nums)
	decreaseNum := 0
	for i := 1; i <= len(nums); i++ {
		if nums[i%numsLen] < nums[i-1] {
			decreaseNum++
			if decreaseNum > 1 {
				return false
			}
		}
	}
	return true
}
