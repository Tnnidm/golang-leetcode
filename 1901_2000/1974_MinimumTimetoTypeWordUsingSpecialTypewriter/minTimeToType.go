package minTimeToType

func minTimeToType(word string) int {
	last := 'a'
	result := 0
	for _, v := range word {
		distance := abs(int(v - last))
		result += min(distance, 26-distance) + 1
		last = v
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
