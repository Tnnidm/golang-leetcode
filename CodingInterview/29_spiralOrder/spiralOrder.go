package spiralOrder

// // 方法一
// func spiralOrder(matrix [][]int) []int {
// 	// 先判断是不是特殊矩阵
// 	row := len(matrix)
// 	if row == 0 {
// 		return []int{}
// 	}
// 	col := len(matrix[0])
// 	if col == 0 {
// 		return []int{}
// 	}
// 	// 按顺时针读取
// 	numberNum := row * col
// 	result := make([]int, numberNum)
// 	index := 0
// 	topLimit, downLimit := 0, row-1
// 	leftLimit, rihgtLimit := 0, col-1
// 	for index < numberNum {
// 		for j := leftLimit; j <= rihgtLimit; j++ {
// 			result[index] = matrix[topLimit][j]
// 			index++
// 		}
// 		if topLimit < downLimit {
// 			topLimit++
// 		} else {
// 			break
// 		}
// 		for i := topLimit; i <= downLimit; i++ {
// 			result[index] = matrix[i][rihgtLimit]
// 			index++
// 		}
// 		if leftLimit < rihgtLimit {
// 			rihgtLimit--
// 		} else {
// 			break
// 		}
// 		for j := rihgtLimit; j >= leftLimit; j-- {
// 			result[index] = matrix[downLimit][j]
// 			index++
// 		}
// 		if topLimit < downLimit {
// 			downLimit--
// 		} else {
// 			break
// 		}
// 		for i := downLimit; i >= topLimit; i-- {
// 			result[index] = matrix[i][leftLimit]
// 			index++
// 		}
// 		if leftLimit < rihgtLimit {
// 			leftLimit++
// 		} else {
// 			break
// 		}
// 	}
// 	return result
// }

// 方法二
func spiralOrder(matrix [][]int) []int {
	// 先判断是不是特殊矩阵
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}
	// 按顺时针读取
	numberNum := row * col
	result := make([]int, numberNum)
	index := 0
	topLimit, downLimit := 0, row-1
	leftLimit, rihgtLimit := 0, col-1
	for index < numberNum {
		for j := leftLimit; index < numberNum && j <= rihgtLimit; j++ {
			result[index] = matrix[topLimit][j]
			index++
		}
		topLimit++
		for i := topLimit; index < numberNum && i <= downLimit; i++ {
			result[index] = matrix[i][rihgtLimit]
			index++
		}
		rihgtLimit--
		for j := rihgtLimit; index < numberNum && j >= leftLimit; j-- {
			result[index] = matrix[downLimit][j]
			index++
		}
		downLimit--
		for i := downLimit; index < numberNum && i >= topLimit; i-- {
			result[index] = matrix[i][leftLimit]
			index++
		}
		leftLimit++
	}
	return result
}
