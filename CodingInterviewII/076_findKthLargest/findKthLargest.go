package findKthLargest

import "math/rand"

func findKthLargest(nums []int, k int) int {
	numsLen := len(nums)
	index := partition(nums, 0, numsLen-1)
	for index != numsLen-k {
		if index > numsLen-k {
			index = partition(nums, 0, index)
		} else {
			index = partition(nums, index, numsLen-1)
		}
	}
	return nums[index]
}

func partition(nums []int, start int, end int) int {
	random := start + rand.Intn(end-start+1)
	swap(nums, random, end)
	small := start - 1
	for i := start; i < end; i++ {
		if nums[i] <= nums[end] {
			small++
			swap(nums, i, small)
		}
	}
	small++
	swap(nums, end, small)
	return small
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
