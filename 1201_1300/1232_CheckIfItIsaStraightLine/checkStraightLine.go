package checkStraightLine

func checkStraightLine(coordinates [][]int) bool {
	nodeNum := len(coordinates)
	if nodeNum == 2 {
		return true
	}
	var x1, x2, y1, y2 int
	for i := 0; i < nodeNum-2; i++ {
		x1 = coordinates[i+2][0] - coordinates[i][0]
		x2 = coordinates[i+1][0] - coordinates[i][0]
		y1 = coordinates[i+2][1] - coordinates[i][1]
		y2 = coordinates[i+1][1] - coordinates[i][1]
		if x1*y2-x2*y1 != 0 {
			return false
		}
	}
	return true
}
