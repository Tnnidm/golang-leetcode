package threeSumClosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	numsLen := len(nums)
	gap := 20000
	result := 0
	sort.Ints(nums)
	var left, right int
	for first := 0; first < numsLen-2; first++ {
		for left, right = first+1, numsLen-1; left < right; {
			if abs(nums[left]+nums[right]+nums[first]-target) < gap {
				gap = abs(nums[left] + nums[right] + nums[first] - target)
				result = nums[left] + nums[right] + nums[first]
			}
			if nums[left]+nums[right]+nums[first] <= target {
				left++
			} else {
				right--
			}
		}
		for first < numsLen-2 && nums[first] == nums[first+1] {
			first++
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
