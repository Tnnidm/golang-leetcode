package isBoomerang

func isBoomerang(points [][]int) bool {
	x0 := points[1][0] - points[0][0]
	y0 := points[1][1] - points[0][1]
	x1 := points[2][0] - points[0][0]
	y1 := points[2][1] - points[0][1]
	if (x0 == 0 && y0 == 0) || (x1 == 0 && y1 == 0) || (x0 == x1 && y0 == y1) {
		return false
	}
	return x0*y1-x1*y0 != 0
}
