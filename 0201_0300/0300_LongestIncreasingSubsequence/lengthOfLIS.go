package lengthOfLIS

// func lengthOfLIS(nums []int) int {
// 	numsLen := len(nums)
// 	dp := make([]int, numsLen)
// 	result := 0
// 	for i := 0; i < numsLen; i++ {
// 		dp[i] = 1
// 		for j := i - 1; j >= 0; j-- {
// 			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
// 				dp[i] = dp[j] + 1
// 			}
// 		}
// 		if result < dp[i] {
// 			result = dp[i]
// 		}
// 	}
// 	return result
// }

// 贪心法+二分
func lengthOfLIS(nums []int) int {
	dp := []int{}
	dpLen := 0
	for _, num := range nums {
		if dpLen == 0 || dp[dpLen-1] < num {
			dp = append(dp, num)
			dpLen++
		} else if dp[dpLen-1] > num {
			replaceIndex := binarySearch(dp, dpLen, num)
			dp[replaceIndex] = num
		}
	}
	return dpLen
}

func binarySearch(dp []int, dpLen int, num int) int {
	left, right := 0, dpLen-1
	for left < right {
		mid := (left + right) >> 1
		if dp[mid] == num {
			return mid
		} else if dp[mid] > num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
