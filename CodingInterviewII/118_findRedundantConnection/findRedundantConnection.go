package findRedundantConnection

func findRedundantConnection(edges [][]int) []int {
	nodeNum := 0
	for _, edge := range edges {
		nodeNum = max(nodeNum, edge[0])
		nodeNum = max(nodeNum, edge[1])
	}
	nodeNum++
	fathers := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		fathers[i] = i
	}
	for _, edge := range edges {
		father1 := getFather(fathers, edge[0])
		father2 := getFather(fathers, edge[1])
		if father1 != father2 {
			fathers[father1] = father2
		} else {
			return edge
		}
	}
	return []int{}
}

func getFather(fathers []int, i int) int {
	if fathers[i] != i {
		return getFather(fathers, fathers[i])
	}
	return fathers[i]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
