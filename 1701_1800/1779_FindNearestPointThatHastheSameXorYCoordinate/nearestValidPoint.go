package nearestValidPoint

func nearestValidPoint(x int, y int, points [][]int) int {
	minManhattanDistance := 10000
	minManhattanDistancePoint := -1
	for i := 0; i < len(points); i++ {
		if x == points[i][0] {
			manhattanDistance := abs(y - points[i][1])
			if manhattanDistance < minManhattanDistance {
				minManhattanDistance = manhattanDistance
				minManhattanDistancePoint = i
			}
			continue
		}
		if y == points[i][1] {
			manhattanDistance := abs(x - points[i][0])
			if manhattanDistance < minManhattanDistance {
				minManhattanDistance = manhattanDistance
				minManhattanDistancePoint = i
			}
		}
	}
	return minManhattanDistancePoint
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
