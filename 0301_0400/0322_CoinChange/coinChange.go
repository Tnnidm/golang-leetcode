package coinChange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
