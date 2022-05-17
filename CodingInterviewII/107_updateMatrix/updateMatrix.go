package updateMatrix

// 使用动态规划解决，先左上到右下，再右下到左上。
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				result[i][j] = 20000
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				result[i][j] = min(result[i][j], result[i-1][j]+1)
			}
			if j-1 >= 0 {
				result[i][j] = min(result[i][j], result[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				result[i][j] = min(result[i][j], result[i+1][j]+1)
			}
			if j+1 < n {
				result[i][j] = min(result[i][j], result[i][j+1]+1)
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// // 使用BFS解决
// func updateMatrix(mat [][]int) [][]int {
// 	m, n := len(mat), len(mat[0])
// 	result := make([][]int, m)
// 	visit := make([][]bool, m)
// 	queue := [][2]int{}
// 	for i := 0; i < m; i++ {
// 		result[i] = make([]int, n)
// 		visit[i] = make([]bool, n)
// 		for j := 0; j < n; j++ {
// 			if mat[i][j] == 1 {
// 				result[i][j] = 20000
// 			} else {
// 				visit[i][j] = true
// 				queue = append(queue, [2]int{i, j})
// 			}
// 		}
// 	}
// 	move := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 	for len(queue) != 0 {
// 		node := queue[0]
// 		queue = queue[1:]
// 		for i := 0; i < 4; i++ {
// 			x := node[0] + move[i][0]
// 			y := node[1] + move[i][1]
// 			if x >= 0 && x < m && y >= 0 && y < n && !visit[x][y] {
// 				result[x][y] = result[node[0]][node[1]] + 1
// 				queue = append(queue, [2]int{x, y})
// 				visit[x][y] = true
// 			}
// 		}
// 	}
// 	return result
// }
