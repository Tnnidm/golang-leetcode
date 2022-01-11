package findGCD

func findGCD(nums []int) int {
	maxNum, minNum := 0, 1000
	for i := range nums {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	for maxNum%minNum != 0 {
		maxNum, minNum = minNum, maxNum%minNum
	}
	return minNum
}
