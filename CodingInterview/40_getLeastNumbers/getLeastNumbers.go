package getLeastNumbers

import "math/rand"

// // 第一种方法是直接排序后取前k
// // 时间复杂度是O(nlogn)，空间复杂度是O(1)，另外还需要修改原数组
// func getLeastNumbers(arr []int, k int) []int {
// 	sort.Ints(arr)
// 	return arr[:k]
// }

// 第二种方法是用partition函数，去寻找第k个index，因为partition后k的左边一定比k小
// 时间复杂度是O(n)，空间复杂度是O(1)，另外还需要修改原数组
// 时间复杂度相比直接排序，节省在k的左边不要求排序
func getLeastNumbers(arr []int, k int) []int {
	arrLen := len(arr)
	if k == 0 || arrLen == 0 {
		return []int{}
	}
	start, end := 0, arrLen-1
	index := partition(arr, start, end)
	for index != k-1 {
		if index < k-1 {
			start = index + 1
		} else {
			end = index - 1
		}
		index = partition(arr, start, end)
	}
	return arr[:k]
}

func partition(arr []int, start int, end int) int {
	if start == end {
		return start
	}
	index := rand.Intn(end-start) + start
	arr[index], arr[end] = arr[end], arr[index]
	small := start - 1
	for index = start; index < end; index++ {
		if arr[index] < arr[end] {
			small++
			if small != index {
				arr[small], arr[index] = arr[index], arr[small]
			}
		}
	}
	small++
	arr[small], arr[end] = arr[end], arr[small]
	return small
}

// // 第三种方法是用一个数组维护一个k的序列，每到一个新的数，就可以二分查找后插入数组
// // 这样的时间复杂度是O(nlogk)，空间复杂度是O(k)，但是不需要修改原数组
// func getLeastNumbers(arr []int, k int) []int {
// 	arrLen := len(arr)
// 	if arrLen == 0 || k == 0 {
// 		return []int{}
// 	}
// 	if arrLen == k {
// 		return arr
// 	}
// 	result := make([]int, k)
// 	for i := 0; i < k; i++ {
// 		result[i] = math.MaxInt32
// 	}
// 	for i := 0; i < arrLen; i++ {
// 		if arr[i] < result[k-1] {
// 			index := binarySearch(result, 0, k-1, arr[i])
// 			for j := k - 1; j > index; j-- {
// 				result[j] = result[j-1]
// 			}
// 			result[index] = arr[i]
// 		}
// 	}
// 	return result
// }

// func binarySearch(result []int, start, end, value int) int {
// 	if start == end {
// 		return start
// 	}
// 	middle := (end-start)>>1 + start
// 	if value == result[middle] {
// 		return middle
// 	} else if value > result[middle] {
// 		return binarySearch(result, middle+1, end, value)
// 	} else {
// 		return binarySearch(result, start, middle, value)
// 	}
// }
