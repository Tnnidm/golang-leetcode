package minCostToMoveChips

func minCostToMoveChips(position []int) int {
	oddNum := 0
	for _, v := range position {
		oddNum += v % 2
	}
	return min(oddNum, len(position)-oddNum)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
