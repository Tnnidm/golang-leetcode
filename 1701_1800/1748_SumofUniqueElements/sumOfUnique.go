package sumOfUnique

import "sort"

func sumOfUnique(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		switch i {
		case 0:
			if nums[i] != nums[i+1] {
				sum = sum + nums[i]
			}
		case len(nums) - 1:
			if nums[i] != nums[i-1] {
				sum = sum + nums[i]
			}
		default:
			if nums[i] != nums[i+1] && nums[i] != nums[i-1] {
				sum = sum + nums[i]
			}
		}
	}
	return sum
}
