package distanceBetweenBusStops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if destination < start {
		start, destination = destination, start
	}
	dis1, dis2 := 0, 0
	for i := 0; i < len(distance); i++ {
		if i >= start && i < destination {
			dis1 += distance[i]
		} else {
			dis2 += distance[i]
		}
	}
	return min(dis1, dis2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
