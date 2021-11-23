package search

func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
