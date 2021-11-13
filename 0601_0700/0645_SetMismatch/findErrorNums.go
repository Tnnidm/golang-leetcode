package findErrorNums

import "sort"

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	result := []int{}
	numsLen := len(nums)
	diff := (1+numsLen)*numsLen/2 - nums[0]
	for i := 1; i < len(nums); i++ {
		diff -= nums[i]
		if nums[i] == nums[i-1] {
			result = append(result, nums[i])
		}
	}
	result = append(result, result[0]+diff)
	return result
}
