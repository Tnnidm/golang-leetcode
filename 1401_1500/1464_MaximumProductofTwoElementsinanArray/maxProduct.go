package maxProduct

func maxProduct(nums []int) int {
	mostMax, secondMax := 0, 0
	for i := range nums {
		if nums[i] > secondMax {
			secondMax = nums[i]
			if secondMax > mostMax {
				secondMax, mostMax = mostMax, secondMax
			}
		}
	}
	return (mostMax - 1) * (secondMax - 1)
}
