package searchInsert

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
