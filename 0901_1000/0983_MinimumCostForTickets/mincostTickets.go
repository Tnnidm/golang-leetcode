package mincostTickets

func mincostTickets(days []int, costs []int) int {
	// 动态规划dp[i]表示第i天到最后一天的最小花费
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+2)
	index := len(days) - 1
	for i := lastDay; i >= 0; i-- {
		if index >= 0 && i == days[index] {
			dp[i] = dp[i+1] + costs[0]
			if i+7 <= lastDay {
				dp[i] = min(dp[i], dp[i+7]+costs[1])
			} else {
				dp[i] = min(dp[i], costs[1])
			}
			if i+30 <= lastDay {
				dp[i] = min(dp[i], dp[i+30]+costs[2])
			} else {
				dp[i] = min(dp[i], costs[2])
			}
			index--
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
