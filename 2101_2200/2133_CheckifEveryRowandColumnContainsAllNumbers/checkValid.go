package checkValid

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	var contains []bool
	// 遍历每一行
	for i := 0; i < n; i++ {
		// 初始化contains
		contains = make([]bool, n)
		for j := 0; j < n; j++ {
			contains[matrix[i][j]-1] = true
		}
		for k := 0; k < n; k++ {
			if contains[k] == false {
				return false
			}
		}
	}
	// 遍历每一列
	for j := 0; j < n; j++ {
		// 初始化contains
		contains = make([]bool, n)
		for i := 0; i < n; i++ {
			contains[matrix[i][j]-1] = true
		}
		for k := 0; k < n; k++ {
			if contains[k] == false {
				return false
			}
		}
	}
	return true
}
