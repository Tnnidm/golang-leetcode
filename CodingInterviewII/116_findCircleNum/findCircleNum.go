package findCircleNum

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for neighbor, connection := range isConnected[node] {
			if !visited[neighbor] && connection == 1 {
				dfs(neighbor)
			}
		}
	}
	result := 0
	for node, beVisited := range visited {
		if !beVisited {
			result++
			dfs(node)
		}
	}
	return result
}
