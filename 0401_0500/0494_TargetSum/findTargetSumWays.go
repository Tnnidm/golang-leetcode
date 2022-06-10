package findTargetSumWays

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target += sum
	if target < 0 || target%2 != 0 {
		return 0
	}
	target >>= 1
	numsLen := len(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < numsLen; i++ {
		for j := target - nums[i]; j >= 0; j-- {
			dp[j+nums[i]] += dp[j]
		}
	}
	return dp[target]
}
