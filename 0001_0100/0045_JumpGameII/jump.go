package jump

// // DP
// func jump(nums []int) int {
// 	numsLen := len(nums)
// 	dp := make([]int, numsLen)
// 	for i := numsLen - 2; i >= 0; i-- {
// 		for j := i + 1; j <= i+nums[i]; j++ {
// 			dp[i] = min(dp[i], dp[j])
// 		}
// 		dp[i]++
// 	}
// 	return dp[0]
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// 贪心
func jump(nums []int) int {
	numsLen := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < numsLen-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
