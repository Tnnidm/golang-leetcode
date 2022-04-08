package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	result := [][]int{}
	for firstIndex := 0; firstIndex < numsLen-2; {
		left, right := firstIndex+1, numsLen-1
		for left < right {
			if nums[left]+nums[right] == -nums[firstIndex] {
				result = append(result, []int{nums[firstIndex], nums[left], nums[right]})
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if nums[left]+nums[right] < -nums[firstIndex] {
				left++
			} else {
				right--
			}
		}
		for firstIndex+1 < numsLen-2 && nums[firstIndex] == nums[firstIndex+1] {
			firstIndex++
		}
		firstIndex++
	}
	return result
}
