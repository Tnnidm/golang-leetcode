package replaceElements

func replaceElements(arr []int) []int {
	maxRight := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > maxRight {
			arr[i], maxRight = maxRight, arr[i]
		} else {
			arr[i] = maxRight
		}
	}
	arr[len(arr)-1] = -1
	return arr
}
