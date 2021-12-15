package minTimeToVisitAllPoints

func minTimeToVisitAllPoints(points [][]int) int {
	counter := 0
	for i := 0; i < len(points)-1; i++ {
		counter = counter + max(abs(points[i+1][0]-points[i][0]), abs(points[i+1][1]-points[i][1]))
	}
	return counter
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
