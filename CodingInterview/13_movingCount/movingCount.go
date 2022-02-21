package movingCount

func digitSum(n int) int {
	sum := 0
	for ; n != 0; n = n / 10 {
		sum += n % 10
	}
	return sum
}

// // 广度优先搜索法
// func movingCount(m int, n int, k int) int {
// 	used := make([][]bool, m)
// 	for i := 0; i < m; i++ {
// 		used[i] = make([]bool, n)
// 	}
// 	count := 1
// 	queue := [][]int{{0, 0}}
// 	used[0][0] = true
// 	for len(queue) != 0 {
// 		i := queue[0][0]
// 		j := queue[0][1]
// 		if i < m-1 && !used[i+1][j] && digitSum(i+1)+digitSum(j) <= k {
// 			queue = append(queue, []int{i + 1, j})
// 			used[i+1][j] = true
// 			count++
// 		}
// 		if j < n-1 && !used[i][j+1] && digitSum(i)+digitSum(j+1) <= k {
// 			queue = append(queue, []int{i, j + 1})
// 			used[i][j+1] = true
// 			count++
// 		}
// 		queue = queue[1:]
// 	}
// 	return count
// }

// 深度优先算法/回溯法
func movingCount(m int, n int, k int) int {
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, &used)
}

func dfs(i, j, m, n, k int, used *[][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || (*used)[i][j] {
		return 0
	}
	if digitSum(i)+digitSum(j) <= k {
		(*used)[i][j] = true
		return 1 + dfs(i+1, j, m, n, k, used) + dfs(i, j+1, m, n, k, used)
	} else {
		return 0
	}
}
