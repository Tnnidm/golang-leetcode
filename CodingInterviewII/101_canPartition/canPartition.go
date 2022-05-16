package canPartition

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum >> 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for j := sum; j > 0; j-- {
			if j-num >= 0 {
				dp[j] = dp[j-num] || dp[j]
			}
		}
	}
	return dp[sum]
}
