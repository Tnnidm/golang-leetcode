package singleNonDuplicate

func singleNonDuplicate(nums []int) int {
	numsLen := len(nums)
	left, right := 0, numsLen>>1
	for left <= right {
		mid := left + (right-left)>>1
		index := mid << 1
		if index == numsLen-1 {
			return nums[index]
		}
		if nums[index] != nums[index+1] {
			if index == 0 || nums[index-2] == nums[index-1] {
				return nums[index]
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
