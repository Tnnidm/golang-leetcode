package longestIncreasingPath

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	pathLen := make([][]int, m)
	for i := 0; i < m; i++ {
		pathLen[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if pathLen[i][j] != 0 {
			return pathLen[i][j]
		}
		pathLen[i][j] = 1
		if i > 0 && matrix[i][j] < matrix[i-1][j] {
			pathLen[i][j] = max(pathLen[i][j], dfs(i-1, j)+1)
		}
		if i < m-1 && matrix[i][j] < matrix[i+1][j] {
			pathLen[i][j] = max(pathLen[i][j], dfs(i+1, j)+1)
		}
		if j > 0 && matrix[i][j] < matrix[i][j-1] {
			pathLen[i][j] = max(pathLen[i][j], dfs(i, j-1)+1)
		}
		if j < n-1 && matrix[i][j] < matrix[i][j+1] {
			pathLen[i][j] = max(pathLen[i][j], dfs(i, j+1)+1)
		}
		return pathLen[i][j]
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = max(result, dfs(i, j))
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
