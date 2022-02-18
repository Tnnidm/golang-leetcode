package typicalSort

// 计数排序是一个数一个桶的桶排序
func countingSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	// 扫描确定桶的个数，也判断能不能用计数排序
	minValue := arr[0]
	maxValue := arr[0]
	for i := 0; i < arrLen; i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	bucket := make([]int, maxValue-minValue+1)
	// 统计各个桶的个数
	for i := 0; i < arrLen; i++ {
		bucket[arr[i]-minValue]++
	}
	// 重写arr
	index := 0
	for i := 0; i < maxValue-minValue+1; i++ {
		for ; bucket[i] > 0; bucket[i]-- {
			arr[index] = i + minValue
			index++
		}
	}
	return arr
}
