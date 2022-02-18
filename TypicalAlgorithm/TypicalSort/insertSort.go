package typicalSort

func insertSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	for i := 1; i < arrLen; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
