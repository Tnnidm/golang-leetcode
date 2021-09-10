package maxProfit

// 执行用时：156 ms, 在所有 Go 提交中击败了19.21%的用户
// 内存消耗：7.9 MB, 在所有 Go 提交中击败了92.49%的用户
func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}
func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen == 1 {
		return 0
	} else {
		profit := 0
		maxprofit := 0
		for i := pricesLen - 2; i >= 0; i-- {
			profit = max(profit, 0) + prices[i+1] - prices[i]
			maxprofit = max(profit, maxprofit)
		}
		return maxprofit
	}
}

// // Other's method
// // 执行用时：120 ms, 在所有 Go 提交中击败了69.09%的用户
// // 内存消耗：8.1 MB, 在所有 Go 提交中击败了82.78%的用户
// func maxProfit(prices []int) int {
// 	minPrice := prices[0]
// 	maxProfit := 0
// 	for i := range prices {
// 		if prices[i] < minPrice {
// 			minPrice = prices[i]
// 		} else {
// 			if prices[i]-minPrice > maxProfit {
// 				maxProfit = prices[i] - minPrice
// 			}
// 		}
// 	}
// 	return maxProfit
// }
