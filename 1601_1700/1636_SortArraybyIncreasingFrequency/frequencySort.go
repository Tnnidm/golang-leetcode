package frequencySort

import "sort"

func frequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] >= nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}
