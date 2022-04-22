package searchRange

func searchRange(nums []int, target int) []int {
	lastIndex := len(nums) - 1
	result := []int{-1, -1}
	var left, right, mid int
	for left, right = 0, lastIndex; left <= right; {
		mid = (left + right) >> 1
		if nums[mid] == target && (mid == 0 || nums[mid-1] < target) {
			result[0] = mid
			break
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	for left, right = 0, lastIndex; left <= right; {
		mid = (left + right) >> 1
		if nums[mid] == target && (mid == lastIndex || nums[mid+1] > target) {
			result[1] = mid
			break
		} else if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
