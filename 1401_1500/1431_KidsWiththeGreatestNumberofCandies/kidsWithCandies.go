package kidsWithCandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	kidNum := len(candies)
	for i := 0; i < kidNum; i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	result := make([]bool, kidNum)
	for i := 0; i < kidNum; i++ {
		result[i] = (candies[i] + extraCandies) >= maxCandies
	}
	return result
}
