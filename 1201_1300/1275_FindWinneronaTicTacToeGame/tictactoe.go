package tictactoe

func tictactoe(moves [][]int) string {
	board := make([][]int, 3)
	for k := range board {
		board[k] = make([]int, 3)
	}
	for k, v := range moves {
		if k%2 == 0 {
			board[v[0]][v[1]]++
		} else {
			board[v[0]][v[1]]--
		}
	}
	l := len(moves)
	r := win(board, moves[l-1][0], moves[l-1][1])
	if r != "" {
		return r
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func win(board [][]int, row, col int) (result string) {
	tRow, tCol, tL2r, tR2l := 0, 0, 0, 0
	for i := 0; i < 3; i++ {
		tRow += board[row][i]
	}
	for i := 0; i < 3; i++ {
		tCol += board[i][col]
	}
	if row == col {
		for i := 0; i < 3; i++ {
			tL2r += board[i][i]
		}
	}
	if row == 2-col {
		for i := 0; i < 3; i++ {
			tR2l += board[i][2-i]
		}
	}
	if tRow == 3 || tCol == 3 || tL2r == 3 || tR2l == 3 {
		return "A"
	}
	if tRow == -3 || tCol == -3 || tL2r == -3 || tR2l == -3 {
		return "B"
	}
	return ""
}
