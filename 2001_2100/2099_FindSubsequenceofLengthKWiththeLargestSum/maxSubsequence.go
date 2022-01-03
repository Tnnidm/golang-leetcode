package maxSubsequence

import "sort"

func maxSubsequence(nums []int, k int) []int {
	numsLen := len(nums)
	index := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return nums[index[i]] > nums[index[j]]
	})
	result := index[:k]
	sort.Ints(result)
	for i := 0; i < k; i++ {
		result[i] = nums[result[i]]
	}
	return result
}
