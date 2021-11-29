package peakIndexInMountainArray

func peakIndexInMountainArray(arr []int) int {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		if arr[i] > arr[i+1] {
			return i
		}
		if arr[j] > arr[j-1] {
			return j
		}
	}
	return -1
}
