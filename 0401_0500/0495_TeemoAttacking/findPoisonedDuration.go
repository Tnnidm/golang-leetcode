package findPoisonedDuration

func findPoisonedDuration(timeSeries []int, duration int) int {
	timeLen := len(timeSeries)
	if timeLen == 0 {
		return 0
	} else {
		result := duration
		for i := 0; i < len(timeSeries)-1; i++ {
			result += min(timeSeries[i+1]-timeSeries[i], duration)
		}
		return result
	}
}

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
