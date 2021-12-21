package createTargetArray

func createTargetArray(nums []int, index []int) []int {
	numLen := len(nums)
	target := make([]int, numLen)
	for i := 0; i < numLen; i++ {
		for j := len(index) - 1; j >= index[i]+1; j-- {
			target[j] = target[j-1]
		}
		target[index[i]] = nums[i]
	}
	return target
}
