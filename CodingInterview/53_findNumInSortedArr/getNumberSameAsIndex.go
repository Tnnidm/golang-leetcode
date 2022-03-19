package findNumInSortedArr

func getNumberSameAsIndex(nums []int) int {
	begin, end := 0, len(nums)-1
	var mid int
	for begin <= end {
		mid = begin + (end-begin)>>1
		if mid == 0 || nums[mid] == mid {
			return mid
		} else if nums[mid] > mid {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}
