package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	result := [][]int{}
	var left, right int
	for first := 0; first < numsLen-3; first++ {
		for second := first + 1; second < numsLen-2; second++ {
			for left, right = second+1, numsLen-1; left < right; {
				if nums[first]+nums[second]+nums[left]+nums[right] > target {
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
				} else {
					if nums[first]+nums[second]+nums[left]+nums[right] == target {
						result = append(result, []int{nums[first], nums[second], nums[left], nums[right]})
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
				}
			}
			for second < numsLen-2 && nums[second] == nums[second+1] {
				second++
			}
		}
		for first < numsLen-3 && nums[first] == nums[first+1] {
			first++
		}
	}
	return result
}
