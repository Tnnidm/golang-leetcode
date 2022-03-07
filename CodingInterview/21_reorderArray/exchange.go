package reorderArray

func exchange(nums []int) []int {
	for begin, end := 0, len(nums)-1; begin < end; begin, end = begin+1, end-1 {
		for nums[begin]&1 == 1 && begin < end {
			begin++
		}
		for nums[end]&1 == 0 && begin < end {
			end--
		}
		nums[begin], nums[end] = nums[end], nums[begin]
	}
	return nums
}
