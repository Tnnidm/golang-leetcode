package maximumProduct

import "sort"

func maximumProduct(nums []int) int {
	numsLen := len(nums)
	sort.Ints(nums)
	return max(nums[numsLen-1]*nums[numsLen-2]*nums[numsLen-3], nums[numsLen-1]*nums[0]*nums[1])
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
