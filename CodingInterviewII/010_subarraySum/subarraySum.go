package subarraySum

func subarraySum(nums []int, k int) int {
	exist := make(map[int]int, len(nums)+1)
	exist[0] = 1
	preSum, result := 0, 0
	for _, num := range nums {
		preSum += num
		if times, ok := exist[preSum-k]; ok {
			result += times
		}
		exist[preSum]++
	}
	return result
}
