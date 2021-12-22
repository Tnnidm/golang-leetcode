package minSubsequence

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	numLen := len(nums)
	lower, higher := 0, numLen
	smallerSum, largerSum := 0, 0
	for lower < higher {
		higher--
		largerSum = largerSum + nums[higher]
		for smallerSum+nums[lower] < largerSum && lower < higher {
			smallerSum = smallerSum + nums[lower]
			lower++
		}
	}
	result := nums[higher:]
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
