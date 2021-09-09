package maxSubArray

func max(num1 int, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

// My method 动态规划
// 执行用时：4 ms, 在所有 Go 提交中击败了96.56%的用户
// 内存消耗：3.4 MB, 在所有 Go 提交中击败了17.49%的用户
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	var maxSum int
	if numsLen == 1 {
		maxSum = nums[0]
	} else {
		dp := make([]int, numsLen)
		dp[0] = nums[0]
		maxSum = nums[0]
		for i := 1; i < numsLen; i++ {
			dp[i] = max(nums[i], dp[i-1]+nums[i])
			maxSum = max(dp[i], maxSum)
		}
	}
	return maxSum
}

// // others' method 分治法
// // 执行用时：108 ms, 在所有 Go 提交中击败了5.04%的用户
// // 内存消耗：3.2 MB, 在所有 Go 提交中击败了30.45%的用户
// func maxSubArray(nums []int) int {
// 	return calc(0, len(nums)-1, nums)
// }

// func calc(l, r int, nums []int) int {
// 	if l == r {
// 		return nums[l]
// 	}
// 	mid := (l + r) >> 1
// 	lmax, lsum, rmax, rsum := nums[mid], 0, nums[mid+1], 0

// 	for i := mid; i >= 0; i-- {
// 		lsum += nums[i]
// 		lmax = max(lmax, lsum)
// 	}

// 	for i := mid + 1; i <= r; i++ {
// 		rsum += nums[i]
// 		rmax = max(rmax, rsum)
// 	}
// 	return max(max(calc(l, mid, nums), calc(mid+1, r, nums)), lmax+rmax)
// }
