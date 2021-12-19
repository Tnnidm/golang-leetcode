package decompressRLElist

func decompressRLElist(nums []int) []int {
	valNum := len(nums) / 2
	decodeNum := 0
	for i := 0; i < valNum; i++ {
		decodeNum += nums[2*i]
	}
	result := make([]int, decodeNum)
	index := 0
	for i := 0; i < valNum; i++ {
		for j := 0; j < nums[2*i]; j++ {
			result[index] = nums[2*i+1]
			index++
		}
	}
	return result
}
