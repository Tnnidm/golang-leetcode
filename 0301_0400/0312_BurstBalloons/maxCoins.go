package maxCoins

func maxCoins(nums []int) int {
	numsLen := len(nums)
	val := make([]int, numsLen+2)
	for i := 1; i <= numsLen; i++ {
		val[i] = nums[i-1]
	}
	val[0] = 1
	val[numsLen+1] = 1
	// dp := make([][]int, numsLen+2)
	// for i := range dp {
	//     dp[i] = make([]int, numsLen+2)
	// }
	dp := make(map[[2]int]int, (numsLen+2)*(numsLen+2)/2)
	for length := 2; length <= numsLen+1; length++ {
		for i := 0; i+length <= numsLen+1; i++ {
			for j := i + 1; j <= i+length-1; j++ {
				dp[[2]int{i, i + length}] = max(dp[[2]int{i, i + length}], dp[[2]int{i, j}]+dp[[2]int{j, i + length}]+val[i]*val[j]*val[i+length])
			}
		}
	}
	return dp[[2]int{0, numsLen + 1}]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
