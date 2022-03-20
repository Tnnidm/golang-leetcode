package maxProfit

func maxProfit(prices []int) int {
	priceLen := len(prices)
	if priceLen < 2 {
		return 0
	}
	buyPrice := prices[0]
	maxProfit := 0
	for i := 1; i < priceLen; i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buyPrice)
		}
	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
