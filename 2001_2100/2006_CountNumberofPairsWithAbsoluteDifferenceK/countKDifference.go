package countKDifference

func countKDifference(nums []int, k int) int {
	numsLen := len(nums)
	numsMap := make(map[int]int, 2*numsLen)
	result := 0
	for i := 0; i < numsLen; i++ {
		result += numsMap[nums[i]]
		numsMap[nums[i]+k]++
		numsMap[nums[i]-k]++
	}
	return result
}
