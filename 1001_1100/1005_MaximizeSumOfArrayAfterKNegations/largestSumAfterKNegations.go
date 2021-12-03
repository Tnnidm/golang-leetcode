package largestSumAfterKNegations

import "sort"

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		} else {
			break
		}
	}
	if k%2 > 0 {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	result := 0
	for i := range nums {
		result += nums[i]
	}
	return result
}
