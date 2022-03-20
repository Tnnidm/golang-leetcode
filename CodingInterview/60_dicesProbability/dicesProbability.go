package dicesProbability

import (
	"math"
)

func dicesProbability(n int) []float64 {
	avaialbleSum := math.Pow(6.0, float64(n))
	lastTimes := make([]float64, 6*n+1)
	currentTimes := make([]float64, 6*n+1)
	for i := 1; i <= 6; i++ {
		lastTimes[i] = 1.0
	}
	for temp := 2; temp <= n; temp++ {
		for i := 0; i <= 6*temp; i++ {
			currentTimes[i] = 0
			for j := max(i-6, 1); j <= i-1; j++ {
				currentTimes[i] += lastTimes[j]
			}
		}
		for i := 0; i <= 6*temp; i++ {
			lastTimes[i] = currentTimes[i]
		}
	}
	lastTimes = lastTimes[n : 6*n+1]
	for i := 0; i < 5*n+1; i++ {
		lastTimes[i] /= avaialbleSum
	}
	return lastTimes
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
