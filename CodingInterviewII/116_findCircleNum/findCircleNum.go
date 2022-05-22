package findCircleNum

// // DFS方法
// func findCircleNum(isConnected [][]int) int {
// 	visited := make([]bool, len(isConnected))
// 	var dfs func(int)
// 	dfs = func(node int) {
// 		visited[node] = true
// 		for neighbor, connection := range isConnected[node] {
// 			if !visited[neighbor] && connection == 1 {
// 				dfs(neighbor)
// 			}
// 		}
// 	}
// 	result := 0
// 	for node, beVisited := range visited {
// 		if !beVisited {
// 			result++
// 			dfs(node)
// 		}
// 	}
// 	return result
// }

// 并查集
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	father := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
	}
	result := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && union(father, i, j) {
				result--
			}
		}
	}
	return result
}

func findFather(father []int, i int) int {
	temp := father[i]
	for temp != father[temp] {
		temp = father[temp]
	}
	father[i] = temp
	return temp
}

func union(father []int, i, j int) bool {
	fatherOfi := findFather(father, i)
	fatherOfj := findFather(father, j)
	if fatherOfi != fatherOfj {
		father[fatherOfi] = fatherOfj
		return true
	}
	return false
}
