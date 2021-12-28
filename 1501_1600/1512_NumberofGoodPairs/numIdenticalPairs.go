package numIdenticalPairs

func numIdenticalPairs(nums []int) int {
	result := 0
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}
	return result
}
