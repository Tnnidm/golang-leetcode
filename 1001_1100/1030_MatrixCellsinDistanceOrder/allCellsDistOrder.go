package main

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	maxDis := max(rCenter, rows-1-rCenter) + max(cCenter, cols-1-cCenter)
	result := make([][]int, rows*cols)
	var rowDis, colDis, index int
	for dis := 0; dis <= maxDis; dis++ {
		for row := max(rCenter-dis, 0); row <= min(rCenter+dis, rows-1); row++ {
			rowDis = abs(row - rCenter)
			colDis = dis - rowDis
			if cCenter-colDis >= 0 {
				result[index] = []int{row, cCenter - colDis}
				index++
			}
			if cCenter+colDis <= cols-1 && colDis != 0 {
				result[index] = []int{row, cCenter + colDis}
				index++
			}
		}
	}
	return result
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	allCellsDistOrder(1, 2, 0, 0)
}
