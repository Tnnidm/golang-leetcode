package typicalSort

func heapSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	// 建堆
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
	// 依次拿掉堆头
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		arrLen--
		heapify(arr, 0, arrLen)
	}
	return arr
}

func heapify(arr []int, i, arrLen int) {
	largestIndex := i
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	if leftIndex < arrLen && arr[leftIndex] > arr[largestIndex] {
		largestIndex = leftIndex
	}
	if rightIndex < arrLen && arr[rightIndex] > arr[largestIndex] {
		largestIndex = rightIndex
	}
	if largestIndex != i {
		arr[i], arr[largestIndex] = arr[largestIndex], arr[i]
		heapify(arr, largestIndex, arrLen)
	}
}
