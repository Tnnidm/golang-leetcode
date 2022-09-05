package findSubarrays

func findSubarrays(nums []int) bool {
	numsLen := len(nums)
	twoSum := make(map[int]bool, numsLen-1)
	temp := nums[0] + nums[1]
	twoSum[temp] = true
	for i := 1; i < numsLen-1; i++ {
		temp += nums[i+1] - nums[i-1]
		if twoSum[temp] {
			return true
		} else {
			twoSum[temp] = true
		}
	}
	return false
}
