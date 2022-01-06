package canFormArray

func canFormArray(arr []int, pieces [][]int) bool {
	pieceNum := len(pieces)
	pieceMap := make(map[int]int, pieceNum)
	for i, piece := range pieces {
		pieceMap[piece[0]] = i
	}
	for index := 0; index < len(arr); {
		if i, ok := pieceMap[arr[index]]; ok {
			for j := 0; j < len(pieces[i]); index, j = index+1, j+1 {
				if arr[index] != pieces[i][j] {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}
