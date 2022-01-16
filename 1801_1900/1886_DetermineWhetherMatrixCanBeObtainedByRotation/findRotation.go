package findRotation

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	same := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				same = false
				break
			}
		}
	}
	if same {
		return true
	}
	same = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[n-1-j][i] != target[i][j] {
				same = false
				break
			}
		}
	}
	if same {
		return true
	}
	same = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[n-1-i][n-1-j] != target[i][j] {
				same = false
				break
			}
		}
	}
	if same {
		return true
	}
	same = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[j][n-1-i] != target[i][j] {
				same = false
				break
			}
		}
	}
	if same {
		return true
	}
	return false
}
