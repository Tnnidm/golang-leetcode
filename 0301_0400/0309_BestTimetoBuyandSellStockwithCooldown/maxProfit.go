package maxProfit

// // 动态规划
// // 时间复杂度O(n^2), 空间复杂度O(n)
// func maxProfit(prices []int) int {
// 	priceNum := len(prices)
// 	totalProfit := make([]int, priceNum+2)
// 	for i := priceNum - 2; i >= 0; i-- {
// 		totalProfit[i] = totalProfit[i+1]
// 		for j := i + 1; j < priceNum; j++ {
// 			if prices[j] > prices[i] {
// 				totalProfit[i] = max(totalProfit[i], prices[j]-prices[i]+totalProfit[j+2])
// 			}
// 		}
// 	}
// 	return totalProfit[0]
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// 动态规划
// O(n)时间复杂度，O(n)空间复杂度
func maxProfit(prices []int) int {
	profit1 := -prices[0] // 目前持有一支股票，对应的「累计最大收益」
	profit2 := 0          // 目前不持有任何股票，并且处于冷冻期中
	profit3 := 0          // 前不持有任何股票，并且不处于冷冻期中
	for _, price := range prices[1:] {
		profit1, profit2, profit3 = max(profit1, profit3-price), max(profit2, profit1+price), max(profit3, profit2)
	}
	return max(profit2, profit3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
