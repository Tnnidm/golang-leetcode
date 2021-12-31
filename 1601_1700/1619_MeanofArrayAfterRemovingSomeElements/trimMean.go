package trimMean

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	arrLen := len(arr)
	var sum int
	for i := arrLen / 20; i < arrLen-arrLen/20; i++ {
		sum = sum + arr[i]
	}
	return float64(sum) / (float64(arrLen) * 0.9)
}
