package findNumInSortedArr

func missingNumber(nums []int) int {
	numsLen := len(nums)
	begin, end := 0, numsLen-1
	var mid int
	for begin != end {
		mid = (begin + end) / 2
		if nums[mid] == mid {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	if nums[begin] == begin {
		return begin + 1
	}
	return begin
}
