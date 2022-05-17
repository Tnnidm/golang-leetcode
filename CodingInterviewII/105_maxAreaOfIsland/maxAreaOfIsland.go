package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	var getArea func(i, j int) int
	getArea = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] == 1 && !visit[i][j] {
			visit[i][j] = true
			return 1 + getArea(i-1, j) + getArea(i+1, j) + getArea(i, j-1) + getArea(i, j+1)
		}
		return 0
	}
	maxArea := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxArea = max(maxArea, getArea(i, j))
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
