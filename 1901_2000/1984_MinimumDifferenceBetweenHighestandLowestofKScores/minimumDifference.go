package minimumDifference

import "sort"

func minimumDifference(nums []int, k int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return 0
	}
	sort.Ints(nums)
	minDiff := 100000
	for i := 0; i < numsLen-k+1; i++ {
		if nums[i+k-1]-nums[i] < minDiff {
			minDiff = nums[i+k-1] - nums[i]
		}
	}
	return minDiff
}
