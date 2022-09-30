package longestSubarray

func longestSubarray(nums []int) int {
	maxNum := nums[0]
	maxCount := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			count = 1
		} else {
			count++
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxCount = count
		} else if nums[i] == maxNum && count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
