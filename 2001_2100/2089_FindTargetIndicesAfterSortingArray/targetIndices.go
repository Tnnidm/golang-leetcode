package targetIndices

func targetIndices(nums []int, target int) []int {
	smallerNum := 0
	targetNum := 0
	for _, v := range nums {
		if v < target {
			smallerNum++
		} else if v == target {
			targetNum++
		}
	}
	if targetNum == 0 {
		return []int{}
	}
	result := make([]int, targetNum)
	for i := 0; i < targetNum; i++ {
		result[i] = smallerNum + i
	}
	return result
}
