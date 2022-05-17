package isBipartite

func isBipartite(graph [][]int) bool {
	nodeNum := len(graph)
	colors := make([]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		colors[i] = -1
	}
	for i := 0; i < nodeNum; i++ {
		if colors[i] == -1 {
			if !setColor(graph, colors, i, 0) {
				return false
			}
		}
	}
	return true
}

// // DFS做法
// func setColor(graph [][]int, colors []int, nodeIndex int, color int) bool {
// 	if colors[nodeIndex] >= 0 {
// 		return colors[nodeIndex] == color
// 	}
// 	colors[nodeIndex] = color
// 	for _, neighbor := range graph[nodeIndex] {
// 		if !setColor(graph, colors, neighbor, 1-color) {
// 			return false
// 		}
// 	}
// 	return true
// }

// BFS做法
func setColor(graph [][]int, colors []int, nodeIndex int, color int) bool {
	queue := []int{nodeIndex}
	colors[nodeIndex] = color
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if colors[neighbor] >= 0 {
				if colors[neighbor] == colors[node] {
					return false
				}
			} else {
				queue = append(queue, neighbor)
				colors[neighbor] = 1 - colors[node]
			}
		}
	}
	return true
}
