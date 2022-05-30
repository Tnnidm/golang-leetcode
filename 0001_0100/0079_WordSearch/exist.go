package exist

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	wordLen := len(word)
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if index == wordLen {
			return true
		}
		if i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] && board[i][j] == word[index] {
			visited[i][j] = true
			result := dfs(i-1, j, index+1) || dfs(i+1, j, index+1) || dfs(i, j-1, index+1) || dfs(i, j+1, index+1)
			visited[i][j] = false
			return result
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
