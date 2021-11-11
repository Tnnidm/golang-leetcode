package findLHS

func findLHS(nums []int) int {
	result := 0
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	for k, v := range count {
		if count[k+1] > 0 && count[k+1]+v > result {
			result = count[k+1] + v
		}
	}
	return result
}
