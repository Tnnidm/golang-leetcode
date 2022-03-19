package findContinuousSequence

func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	left, right := 1, 2
	sum := 3
	for left <= target>>1 {
		if sum == target {
			oneResult := make([]int, right-left+1)
			for i := left; i <= right; i++ {
				oneResult[i-left] = i
			}
			result = append(result, oneResult)
			right++
			sum += right
		} else if sum < target {
			right++
			sum += right
		} else {
			sum -= left
			left++
		}
	}
	return result
}
