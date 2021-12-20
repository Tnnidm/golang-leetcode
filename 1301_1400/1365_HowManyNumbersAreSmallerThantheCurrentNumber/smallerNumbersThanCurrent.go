package smallerNumbersThanCurrent

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		result[i] = nums[i]
	}
	sort.Ints(nums)
	indexMap := make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			indexMap[nums[i]] = i
		}
	}
	for i := 0; i < numsLen; i++ {
		result[i] = indexMap[result[i]]
	}
	return result
}
