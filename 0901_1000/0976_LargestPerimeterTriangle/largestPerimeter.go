package largestPerimeter

import "sort"

func largestPerimeter(nums []int) int {
	numsLen := len(nums)
	sort.Ints(nums)
	for i := numsLen - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
