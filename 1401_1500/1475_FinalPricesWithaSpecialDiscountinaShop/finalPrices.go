package finalPrices

func finalPrices(prices []int) []int {
	priceLen := len(prices)
	for i := 0; i < priceLen-1; i++ {
		for j := i + 1; j < priceLen; j++ {
			if prices[j] <= prices[i] {
				prices[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return prices
}
