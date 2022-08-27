package movingCount

func digitalSum(num int) int {
	result := 0
	for ; num != 0; num /= 10 {
		result += num % 10
	}
	return result
}

// 深度优先解法
func movingCount(m int, n int, k int) int {
	arrived := make([][]bool, m)
	for i := range arrived {
		arrived[i] = make([]bool, n)
	}
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i >= m || j >= n || arrived[i][j] || digitalSum(i)+digitalSum(j) > k {
			return 0
		}
		arrived[i][j] = true
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}
	return dfs(0, 0)
}

// // 广度优先解法
// func movingCount(m int, n int, k int) int {
// 	arrived := make([][]bool, m)
// 	for i := range arrived {
// 		arrived[i] = make([]bool, n)
// 	}
// 	queue := [][2]int{{0, 0}}
// 	result := 0
// 	for len(queue) != 0 {
// 		i, j := queue[0][0], queue[0][1]
// 		queue = queue[1:]
// 		result++
// 		arrived[i][j] = true
// 		if i+1 < m && !arrived[i+1][j] && digitalSum(i+1)+digitalSum(j) <= k {
// 			queue = append(queue, [2]int{i + 1, j})
// 		}
// 		if j+1 < n && !arrived[i][j+1] && digitalSum(i)+digitalSum(j+1) <= k {
// 			queue = append(queue, [2]int{i, j + 1})
// 		}
// 	}
// 	return result
// }
