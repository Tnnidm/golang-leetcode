package maximumRows

func maximumRows(mat [][]int, cols int) int {
	m, n := len(mat), len(mat[0])
	transferMat := make([]int, m)
	for i := range transferMat {
		transferMat[i] = transfer(mat[i])
	}
	// 被覆盖就是col表示的那个数&每一行的数等于0
	chooseCols := make([]int, n)
	for i := range chooseCols {
		chooseCols[i] = 1
	}
	// 选出cols列变成0
	result := 0
	var dfs func(index int, remainCols int)
	dfs = func(index int, remainCols int) {
		if remainCols == 0 {
			temp := 0
			chooseColsToNum := transfer(chooseCols)
			for i := range transferMat {
				if chooseColsToNum&transferMat[i] == 0 {
					temp++
				}
			}
			if temp > result {
				result = temp
			}
		} else if index < n {
			dfs(index+1, remainCols)
			chooseCols[index] = 0
			dfs(index+1, remainCols-1)
			chooseCols[index] = 1
		}
	}
	dfs(0, cols)
	return result
}

func transfer(num []int) int {
	result := 0
	for i := range num {
		result += num[i] * 1 << i
	}
	return result
}
