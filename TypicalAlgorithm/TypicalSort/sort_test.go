package typicalSort

import (
	"reflect"
	"testing"
)

func sort(arr []int) []int {
	// arr = quickSort(arr)    // 快速排序
	// arr = bubbleSort(arr)   // 冒泡排序
	// arr = selectSort(arr)   // 选择排序
	// arr = insertSort(arr)   // 插入排序
	// arr = mergeSort(arr)    // 归并排序
	arr = heapSort(arr) // 堆排序
	// arr = shellSort(arr)    // 希尔排序
	// arr = countingSort(arr) // 计数排序
	// arr = bucketSort(arr)   // 桶排序
	// arr = radixSort(arr)    // 基数排序
	return arr
}

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{3, 4, 2, 1, 5}, []int{1, 2, 3, 4, 5}},
		{"case 2", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case 3", []int{1}, []int{1}},
		{"case 4", []int{}, []int{}},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := sort(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
