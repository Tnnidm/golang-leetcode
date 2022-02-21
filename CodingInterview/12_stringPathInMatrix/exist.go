package exist

func exist(board [][]byte, word string) bool {
	rowNum := len(board)
	colNum := len(board[0])
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if existInMatrix(i, j, board, word, rowNum, colNum) {
				return true
			}
		}
	}
	return false
}
func existInMatrix(i int, j int, board [][]byte, word string, rowNum int, colNum int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= rowNum || j < 0 || j >= colNum {
		return false
	}
	if board[i][j] == word[0] {
		temp := board[i][j]
		board[i][j] = ' '
		if existInMatrix(i-1, j, board, word[1:], rowNum, colNum) ||
			existInMatrix(i+1, j, board, word[1:], rowNum, colNum) ||
			existInMatrix(i, j-1, board, word[1:], rowNum, colNum) ||
			existInMatrix(i, j+1, board, word[1:], rowNum, colNum) {
			return true
		}
		board[i][j] = temp
	}
	return false
}
