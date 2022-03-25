package threeSum

import "sort"

// 排序以后双指针
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	sort.Ints(nums)
	var left, right int
	for target := 0; target < numsLen-2; target++ {
		for left, right = target+1, numsLen-1; left < right; {
			if nums[left]+nums[right]+nums[target] == 0 {
				result = append(result, []int{nums[target], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if nums[left]+nums[right]+nums[target] < 0 {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
		for target < numsLen-2 && nums[target] == nums[target+1] {
			target++
		}
	}
	return result
}
