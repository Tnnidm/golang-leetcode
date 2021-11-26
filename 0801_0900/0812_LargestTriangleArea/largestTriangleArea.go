package largestTriangleArea

func largestTriangleArea(points [][]int) float64 {
	var res int
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				x1, y1, x2, y2, x3, y3 := points[i][0], points[i][1], points[j][0],
					points[j][1], points[k][0], points[k][1]
				s := x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)
				s = abs(s)
				res = max(s, res)
			}
		}
	}
	return float64(res) / float64(2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
