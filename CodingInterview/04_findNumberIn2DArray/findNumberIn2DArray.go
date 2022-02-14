package findNumberIn2DArray

// 时间复杂度O(n+m)，空间复杂度O(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	row, col := 0, m-1
	for {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target && col > 0 {
			col--
			continue
		}
		if matrix[row][col] < target && row < n-1 {
			row++
			continue
		}
		break
	}
	return false
}
