package numWaterBottles

func numWaterBottles(numBottles int, numExchange int) int {
	result := numBottles
	for numBottles >= numExchange {
		numBottles = numBottles - numExchange
		result++
		numBottles++
	}
	return result
}
