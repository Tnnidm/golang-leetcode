package findNumInSortedArr

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	begin := searchBegin(&nums, target, 0, numsLen-1, numsLen)
	if nums[begin] != target {
		return 0
	}
	return searchEnd(&nums, target, 0, numsLen-1, numsLen) - begin + 1
}

func searchBegin(nums *[]int, target, begin, end, numsLen int) int {
	if begin == end {
		return begin
	}
	mid := begin + (end-begin)>>1
	if (*nums)[mid] < target {
		if mid == numsLen-1 {
			return mid
		}
		return searchBegin(nums, target, mid+1, end, numsLen)
	} else if (*nums)[mid] > target {
		if mid == 0 {
			return mid
		}
		return searchBegin(nums, target, begin, mid-1, numsLen)
	} else {
		if mid == 0 || (mid > 0 && (*nums)[mid-1] != target) {
			return mid
		} else {
			return searchBegin(nums, target, begin, mid-1, numsLen)
		}
	}
}

func searchEnd(nums *[]int, target, begin, end, numsLen int) int {
	if begin == end {
		return end
	}
	mid := begin + (end-begin)>>1
	if (*nums)[mid] < target {
		if mid == numsLen-1 {
			return mid
		}
		return searchEnd(nums, target, mid+1, end, numsLen)
	} else if (*nums)[mid] > target {
		if mid == 0 {
			return mid
		}
		return searchEnd(nums, target, begin, mid-1, numsLen)
	} else {
		if mid == len(*nums)-1 || (mid < len(*nums)-1 && (*nums)[mid+1] != target) {
			return mid
		} else {
			return searchEnd(nums, target, mid+1, end, numsLen)
		}
	}
}
