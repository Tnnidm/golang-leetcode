package surfaceArea

func surfaceArea(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	result := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			//  计算每一叠贡献了多少表面积
			// 上侧面
			if i == 0 {
				result += grid[i][j]
			} else {
				result += relu(grid[i][j] - grid[i-1][j])
			}
			// 下侧面
			if i == row-1 {
				result += grid[i][j]
			} else {
				result += relu(grid[i][j] - grid[i+1][j])
			}
			// 左侧面
			if j == 0 {
				result += grid[i][j]
			} else {
				result += relu(grid[i][j] - grid[i][j-1])
			}
			// 右侧面
			if j == col-1 {
				result += grid[i][j]
			} else {
				result += relu(grid[i][j] - grid[i][j+1])
			}
			// 顶面和底面
			if grid[i][j] != 0 {
				result += 2
			}
		}
	}
	return result
}

func relu(x int) int {
	if x <= 0 {
		return 0
	}
	return x
}
