package numRookCaptures

func numRookCaptures(board [][]byte) int {
	var rowR, colR int
	found := false
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'R' {
				rowR = i
				colR = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	result := 0
	for i := rowR + 1; ; i++ {
		if i >= len(board) || board[i][colR] == 'B' {
			break
		}
		if board[i][colR] == 'p' {
			result++
			break
		}
	}
	for i := rowR - 1; ; i-- {
		if i < 0 || board[i][colR] == 'B' {
			break
		}
		if board[i][colR] == 'p' {
			result++
			break
		}
	}
	for j := colR + 1; ; j++ {
		if j >= len(board[0]) || board[rowR][j] == 'B' {
			break
		}
		if board[rowR][j] == 'p' {
			result++
			break
		}
	}
	for j := colR - 1; ; j-- {
		if j < 0 || board[rowR][j] == 'B' {
			break
		}
		if board[rowR][j] == 'p' {
			result++
			break
		}
	}
	return result
}
