package hammingDistance

func hammingDistance(x int, y int) int {
	hamming := 0
	for x != 0 || y != 0 {
		hamming += abs(x%2 - y%2)
		x /= 2
		y /= 2
	}
	return hamming
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
