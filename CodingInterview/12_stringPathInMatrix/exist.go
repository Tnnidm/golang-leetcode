package exist

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	wordLen := len(word)
	searched := make([][]bool, m)
	for i := range searched {
		searched[i] = make([]bool, n)
	}
	var dfs func(i int, j int, index int) bool
	dfs = func(i int, j int, index int) bool {
		if index == wordLen {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || searched[i][j] {
			return false
		}
		if board[i][j] == word[index] {
			searched[i][j] = true
			if dfs(i-1, j, index+1) || dfs(i+1, j, index+1) || dfs(i, j-1, index+1) || dfs(i, j+1, index+1) {
				return true
			}
			searched[i][j] = false
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
