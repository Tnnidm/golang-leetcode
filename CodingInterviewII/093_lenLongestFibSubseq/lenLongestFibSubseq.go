package lenLongestFibSubseq

func lenLongestFibSubseq(arr []int) int {
	arrLen := len(arr)
	numMap := make(map[int]int, arrLen)
	for i := 0; i < arrLen; i++ {
		numMap[arr[i]] = i
	}
	dp := make(map[int]int, arrLen)
	result := 0
	for i := 2; i < arrLen; i++ {
		for j := 0; j < i-1 && arr[i] > arr[j]<<1; j++ {
			if index, exist := numMap[arr[i]-arr[j]]; exist {
				dp[i*arrLen+index] = max(dp[index*arrLen+j]+1, 3)
				result = max(result, dp[i*arrLen+index])
			}
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
