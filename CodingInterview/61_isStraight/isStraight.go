package isStraight

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	index := 0
	jokerCount := 0
	for nums[index] == 0 {
		jokerCount++
		index++
	}
	for index < 4 {
		if nums[index] == nums[index+1] {
			return false
		}
		jokerCount -= nums[index+1] - nums[index] - 1
		index++
	}
	return jokerCount >= 0
}
