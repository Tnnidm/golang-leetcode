package nextPermutation

func nextPermutation(nums []int) {
	numsLen := len(nums)
	var smallerIndex int
	for smallerIndex = numsLen - 2; smallerIndex >= 0; smallerIndex-- {
		if nums[smallerIndex] < nums[smallerIndex+1] {
			break
		}
	}
	if smallerIndex != -1 {
		for i := numsLen - 1; i > smallerIndex; i-- {
			if nums[i] > nums[smallerIndex] {
				nums[i], nums[smallerIndex] = nums[smallerIndex], nums[i]
				break
			}
		}
	}
	for i := 1; smallerIndex+i < numsLen-i; i++ {
		nums[smallerIndex+i], nums[numsLen-i] = nums[numsLen-i], nums[smallerIndex+i]
	}
}
