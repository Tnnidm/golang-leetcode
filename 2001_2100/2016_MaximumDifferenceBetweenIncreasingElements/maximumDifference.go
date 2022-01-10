package maximumDifference

func maximumDifference(nums []int) int {
	maxDiff := 0
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[j]-nums[i] > maxDiff {
				maxDiff = nums[j] - nums[i]
			}
		}
	}
	if maxDiff == 0 {
		return -1
	}
	return maxDiff
}
