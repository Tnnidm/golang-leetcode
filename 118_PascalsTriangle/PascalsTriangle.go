package PascalsTriangle

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了50.27%的用户
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
	}
	if numRows >= 3 {
		for row := 2; row <= numRows-1; row++ {
			for column := 1; column <= row-1; column++ {
				result[row][column] = result[row-1][column-1] + result[row-1][column]
			}
		}
	}
	return result
}
