package sumSubRegion

type NumMatrix struct {
	preSumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	nm.preSumMatrix = make([][]int, rowLen)
	nm.preSumMatrix[0] = make([]int, colLen)
	var rowSum int
	for j := 0; j < colLen; j++ {
		rowSum += matrix[0][j]
		nm.preSumMatrix[0][j] = rowSum
	}
	for i := 1; i < rowLen; i++ {
		nm.preSumMatrix[i] = make([]int, colLen)
		rowSum = 0
		for j := 0; j < colLen; j++ {
			rowSum += matrix[i][j]
			nm.preSumMatrix[i][j] = rowSum + nm.preSumMatrix[i-1][j]
		}
	}
	return nm
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.PreSumRegion(row2, col2) - nm.PreSumRegion(row2, col1-1) - nm.PreSumRegion(row1-1, col2) + nm.PreSumRegion(row1-1, col1-1)
}

func (nm *NumMatrix) PreSumRegion(row int, col int) int {
	if row == -1 || col == -1 {
		return 0
	}
	return nm.preSumMatrix[row][col]
}
