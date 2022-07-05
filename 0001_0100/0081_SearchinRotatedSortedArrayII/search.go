package search

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == nums[right] {
			if nums[left] == target {
				return true
			} else {
				left++
				right--
				continue
			}
		}
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] > target && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
