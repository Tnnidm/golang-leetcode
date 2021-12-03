package minDeletionSize

func minDeletionSize(strs []string) int {
	strNum := len(strs)
	strLen := len(strs[0])
	counter := 0
	for col := 0; col < strLen; col++ {
		for row := 0; row < strNum-1; row++ {
			if int(strs[row+1][col])-int(strs[row][col]) < 0 {
				counter++
				break
			}
		}
	}
	return counter
}
