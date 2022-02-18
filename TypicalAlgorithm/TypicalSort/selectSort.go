package typicalSort

func selectSort(arr []int) []int {
	arrLen := len(arr)
	for i := arrLen - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
