package canPartition

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum >>= 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := sum; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
	}
	return dp[sum]
}
