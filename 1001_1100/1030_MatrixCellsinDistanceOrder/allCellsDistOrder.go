package main

// 直接遍历打印方法
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

// // 自定义排序方法
// func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
// 	result := make([][]int, rows*cols)
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			result[i*cols+j] = []int{i, j}
// 		}
// 	}
// 	sort.Slice(result, func(i, j int) bool {
// 		return abs(result[i][0]-rCenter)+abs(result[i][1]-cCenter) < abs(result[j][0]-rCenter)+abs(result[j][1]-cCenter)
// 	})
// 	return result
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// // 桶排序
// func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
// 	result := map[int][][]int{}
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			result[distance(i, j, rCenter, cCenter)] = append(result[distance(i, j, rCenter, cCenter)], []int{i, j})
// 		}
// 	}
// 	ans := [][]int{}
// 	for i := 0; i <= len(result); i++ {
// 		ans = append(ans, result[i]...)
// 	}
// 	return ans
// }

// func distance(i, j, rCenter, cCenter int) int {
// 	return abs(i-rCenter) + abs(j-rCenter)
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

func main() {
	allCellsDistOrder(1, 2, 0, 0)
}
