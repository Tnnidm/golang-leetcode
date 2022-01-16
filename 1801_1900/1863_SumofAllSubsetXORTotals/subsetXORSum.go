package subsetXORSum

func subsetXORSum(nums []int) int {
	numLen := len(nums)
	if numLen == 1 {
		return nums[0]
	}
	subsetNum := twoPower(numLen - 1)
	subSetXorSum := make([]int, subsetNum)
	subSetXorSum[1] = nums[0]
	before := 2
	result := nums[0]
	for i := 1; i < numLen-1; i++ {
		for j := 0; j < before; j++ {
			subSetXorSum[before+j] = nums[i] ^ subSetXorSum[j]
			result += subSetXorSum[before+j]
		}
		before *= 2
	}
	for i := 0; i < subsetNum; i++ {
		result += nums[numLen-1] ^ subSetXorSum[i]
	}
	return result
}

func twoPower(x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= 2
	}
	return result
}
