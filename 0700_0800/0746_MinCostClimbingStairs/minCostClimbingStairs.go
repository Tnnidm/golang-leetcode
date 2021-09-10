package minCostClimbingStairs

// 执行用时：4 ms, 在所有 Go 提交中击败了84.83%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了99.89%的用户
func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	if costLen == 2 {
		return min(cost[0], cost[1])
	} else {
		reach_step_minus_2 := 0
		reach_step_minus_1 := 0
		costSum := 0
		for i := 2; i <= costLen; i++ {
			costSum = min(reach_step_minus_2+cost[i-2], reach_step_minus_1+cost[i-1])
			reach_step_minus_2 = reach_step_minus_1
			reach_step_minus_1 = costSum
		}
		return costSum
	}
}

func min(n1 int, n2 int) int {
	if n1 <= n2 {
		return n1
	} else {
		return n2
	}
}
