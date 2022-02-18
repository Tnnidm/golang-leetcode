package typicalSort

func bubbleSort(arr []int) []int {
	arrLen := len(arr)
	swapFlag := true
	for i := arrLen - 1; i > 0 && swapFlag; i-- {
		swapFlag = false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapFlag = true
			}
		}
	}
	return arr
}
