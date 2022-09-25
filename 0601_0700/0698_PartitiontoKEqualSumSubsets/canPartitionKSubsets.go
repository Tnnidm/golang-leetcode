package canPartitionKSubsets

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	numsLen := len(nums)
	if nums[numsLen-1] > target {
		return false
	}
	dp := make([]bool, 1<<numsLen)
	dp[0] = true
	choosingSum := make([]int, 1<<numsLen)

	for i, reachable := range dp {
		if !reachable {
			continue
		}
		for j, num := range nums {
			if choosingSum[i]+num > target {
				break
			}
			if i&(1<<j) == 0 {
				next := i | (1 << j)
				if !dp[next] {
					choosingSum[next] = (choosingSum[i] + num) % target
					dp[next] = true
				}
			}
		}
	}
	return dp[1<<numsLen-1]
}
