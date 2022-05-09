package minCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	var costThisStair, costLowerStair, costLowerLowerStair int
	for i := 2; i <= costLen; i++ {
		costThisStair = min(costLowerLowerStair+cost[i-2], costLowerStair+cost[i-1])
		costLowerLowerStair = costLowerStair
		costLowerStair = costThisStair
	}
	return costThisStair
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
