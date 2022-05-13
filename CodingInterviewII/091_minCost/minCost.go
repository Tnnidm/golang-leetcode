package minCost

func minCost(costs [][]int) int {
	var minCostWithLastColor [3]int
	for i := 0; i < len(costs); i++ {
		minCostWithLastColor[0], minCostWithLastColor[1], minCostWithLastColor[2] =
			min(minCostWithLastColor[1], minCostWithLastColor[2])+costs[i][0],
			min(minCostWithLastColor[0], minCostWithLastColor[2])+costs[i][1],
			min(minCostWithLastColor[0], minCostWithLastColor[1])+costs[i][2]
	}
	return min(min(minCostWithLastColor[0], minCostWithLastColor[1]), minCostWithLastColor[2])
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
