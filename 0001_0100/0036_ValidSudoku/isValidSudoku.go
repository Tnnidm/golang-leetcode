package isValidSudoku

func isValidSudoku(board [][]byte) bool {
	// 检查每一行
	for i := 0; i < 9; i++ {
		var exist [9]bool
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if exist[board[i][j]-'0'] {
					return false
				}
				exist[board[i][j]-'0'] = true
			}
		}
	}
	// 检查每一列
	for i := 0; i < 9; i++ {
		var exist [9]bool
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if exist[board[j][i]-'0'] {
					return false
				}
				exist[board[j][i]-'0'] = true
			}
		}
	}
	// 检查每个九宫格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var exist [9]bool
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if board[3*i+m][3*j+n] != '.' {
						if exist[board[3*i+m][3*j+n]-'0'] {
							return false
						}
						exist[board[3*i+m][3*j+n]-'0'] = true
					}
				}
			}
		}
	}
	return true
}
