package subarraySum

func subarraySum(nums []int, k int) int {
	sum := 0
	numSet := make(map[int]int, len(nums)+1)
	numSet[0] = 1
	ans := 0
	for _, num := range nums {
		sum += num
		if SumNum, ok := numSet[sum-k]; ok {
			ans += SumNum
		}
		numSet[sum]++
	}
	return ans
}
