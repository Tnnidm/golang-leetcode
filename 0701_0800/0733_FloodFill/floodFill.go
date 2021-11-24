package floodFill

// func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
// 	preColor := image[sr][sc]
// 	if preColor == newColor {
// 		return image
// 	}
// 	image[sr][sc] = newColor
// 	rowLen := len(image)
// 	colLen := len(image[0])
// 	stack := []int{sr*colLen + sc}
// 	for len(stack) != 0 {
// 		node := stack[0]
// 		stack = stack[1:]
// 		nodeRow := node / colLen
// 		nodeCol := node % colLen
// 		if nodeRow < rowLen-1 {
// 			if image[nodeRow+1][nodeCol] == preColor {
// 				stack = append(stack, node+colLen)
// 				image[nodeRow+1][nodeCol] = newColor
// 			}
// 		}
// 		if nodeRow > 0 {
// 			if image[nodeRow-1][nodeCol] == preColor {
// 				stack = append(stack, node-colLen)
// 				image[nodeRow-1][nodeCol] = newColor
// 			}
// 		}
// 		if nodeCol < colLen-1 {
// 			if image[nodeRow][nodeCol+1] == preColor {
// 				stack = append(stack, node+1)
// 				image[nodeRow][nodeCol+1] = newColor
// 			}

// 		}
// 		if nodeCol > 0 {
// 			if image[nodeRow][nodeCol-1] == preColor {
// 				stack = append(stack, node-1)
// 				image[nodeRow][nodeCol-1] = newColor
// 			}
// 		}
// 	}
// 	return image
// }

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return [][]int{}
	}

	ans := image[sr][sc]
	if ans == 1 && newColor == 1 {
		return image
	}
	dfs(image, sr, sc, ans, newColor)
	return image
}

func dfs(image [][]int, sr int, sc int, ans int, newColor int) {
	if image[sr][sc] == ans {
		image[sr][sc] = newColor
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, direction := range directions {
			srx := sr + direction[0]
			scy := sc + direction[1]
			if srx >= 0 && srx < len(image) && scy >= 0 && scy < len(image[0]) {
				dfs(image, srx, scy, ans, newColor)
			}
		}
	}
}
