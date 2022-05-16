package findTargetSumWays

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	sum = (target + sum) >> 1
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[sum]
}
