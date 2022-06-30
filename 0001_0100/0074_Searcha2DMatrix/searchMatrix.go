package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	upper, lower := 0, row-1
	for upper < lower {
		mid := (upper + lower) >> 1
		if matrix[mid][col-1] == target {
			return true
		} else if matrix[mid][col-1] > target {
			lower = mid
		} else {
			upper = mid + 1
		}
	}
	left, right := 0, col-1
	for left < right {
		mid := (left + right) >> 1
		if matrix[upper][mid] == target {
			return true
		} else if matrix[upper][mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return matrix[upper][right] == target
}
