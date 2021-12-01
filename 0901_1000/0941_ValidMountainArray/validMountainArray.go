package validMountainArray

func validMountainArray(arr []int) bool {
	arrLen := len(arr)
	if arrLen < 3 {
		return false
	}
	i := 0
	for i < arrLen-1 && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == arrLen-1 {
		return false
	}
	for i < arrLen-1 && arr[i] > arr[i+1] {
		i++
	}
	return i == len(arr)-1
}
