package allPathsSourceTarget

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	if n == 1 {
		return [][]int{{0}}
	}
	result := [][]int{}
	combination := []int{}
	var dfs func(node int)
	dfs = func(node int) {
		combination = append(combination, node)
		if node == n-1 {
			result = append(result, append([]int{}, combination...))
		} else {
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}
		combination = combination[:len(combination)-1]
	}
	dfs(0)
	return result
}
