package largestAltitude

func largestAltitude(gain []int) int {
	altitude, maxAltitude := 0, 0
	for _, v := range gain {
		altitude = altitude + v
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}
	return maxAltitude
}
