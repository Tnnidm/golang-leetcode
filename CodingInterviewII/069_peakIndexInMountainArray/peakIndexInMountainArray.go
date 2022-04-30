package peakIndexInMountainArray

func peakIndexInMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for left <= right {
		mid := left + (right-left)>>1
		if arr[mid] > arr[mid-1] {
			if arr[mid] > arr[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}
