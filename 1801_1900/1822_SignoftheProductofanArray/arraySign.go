package arraySign

func arraySign(nums []int) int {
	result := 1
	for i := range nums {
		if nums[i] == 0 {
			result = 0
			break
		} else if nums[i] < 0 {
			result = -result
		}
	}
	return result
}
