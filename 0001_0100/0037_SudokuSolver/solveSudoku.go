package solveSudoku

func solveSudoku(board [][]byte) {
	rows := make([][9]bool, 9)
	cols := make([][9]bool, 9)
	blocks := make([][][9]bool, 3)
	for i := range blocks {
		blocks[i] = make([][9]bool, 3)
	}
	queue := [][]int{}
	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				rows[i][board[i][j]-'1'] = true
				cols[j][board[i][j]-'1'] = true
				blocks[i/3][j/3][board[i][j]-'1'] = true
			} else {
				queue = append(queue, []int{i, j})
			}
		}
	}
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(queue) {
			return true
		}
		row, col := queue[index][0], queue[index][1]
		for i := range rows[queue[index][0]] {
			if !rows[row][i] && !cols[col][i] && !blocks[row/3][col/3][i] {
				rows[row][i] = true
				cols[col][i] = true
				blocks[row/3][col/3][i] = true
				board[row][col] = byte(int8(i) + '1')
				if dfs(index + 1) {
					return true
				}
				rows[row][i] = false
				cols[col][i] = false
				blocks[row/3][col/3][i] = false
			}
		}
		return false
	}
	dfs(0)
}
