package searchMatrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	upper, lower := 0, m-1
	left, right := 0, n-1
	for upper <= lower && left <= right {
		if upper <= lower && matrix[lower][left] > target {
			lower--
			continue
		}
		if left <= right && matrix[lower][left] < target {
			left++
			continue
		}
		if upper <= lower && matrix[lower][left] < target {
			upper++
			continue
		}
		if left <= right && matrix[lower][left] > target {
			right--
			continue
		}
		return true
	}
	return false
}
