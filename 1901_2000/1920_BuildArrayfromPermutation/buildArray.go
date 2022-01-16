package buildArray

func buildArray(nums []int) []int {
	numsLen := len(nums)
	ans := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
