package numIslands

func numIslands(grid [][]byte) (res int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i >= m || j >= n || i < 0 || j < 0 || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
