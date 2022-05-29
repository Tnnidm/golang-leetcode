package canJump

func canJump(nums []int) bool {
	end := 0
	maxPosition := 0
	lastIndex := len(nums) - 1
	for i := 0; i <= lastIndex; i++ {
		if i+nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		if i == end {
			end = maxPosition
			if end >= lastIndex {
				return true
			}
		}
	}
	return false
}
