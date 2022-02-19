package typicalSort

func shellSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	gap := 1
	for gap < arrLen/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < arrLen; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}
