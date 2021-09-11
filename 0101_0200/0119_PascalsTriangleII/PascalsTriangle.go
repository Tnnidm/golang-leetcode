package PascalsTriangleII

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了84.37%的用户
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	} else {
		result := make([]int, rowIndex+1)
		result[0] = 1
		result[1] = 1
		for row := 2; row <= rowIndex; row++ {
			result[row] = 1
			for colume := row - 1; colume >= 1; colume-- {
				result[colume] = result[colume] + result[colume-1]
			}
			print(result)
		}
		return result
	}
}
