package typicalSort

import "sort"

func bucketSort(arr []int) []int {
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
	// 假设桶的数量为5
	bucketNum := 5
	bucket := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucket[i] = []int{}
	}
	var gap = float64(maxValue-minValue) / float64(bucketNum-1)
	for i := 0; i < arrLen; i++ {
		bucketIndex := int(float64(arr[i]-minValue) / gap)
		bucket[bucketIndex] = append(bucket[bucketIndex], arr[i])
	}
	for i := 0; i < bucketNum; i++ {
		sort.Ints(bucket[i])
	}
	index := 0
	for i := 0; i < bucketNum; i++ {
		for j := 0; j < len(bucket[i]); j++ {
			arr[index] = bucket[i][j]
			index++
		}
	}
	return arr
}
