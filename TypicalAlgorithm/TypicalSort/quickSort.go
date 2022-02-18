package typicalSort

func quickSort(arr []int) []int {
	QuickSort(&arr, 0, len(arr)-1)
	return arr
}

func QuickSort(arr *[]int, l int, r int) {
	if l < r {
		small := l - 1
		for i := l; i < r; i++ {
			if (*arr)[i] <= (*arr)[r] {
				small++
				if small != i {
					(*arr)[i], (*arr)[small] = (*arr)[small], (*arr)[i]
				}
			}
		}
		small++
		(*arr)[r], (*arr)[small] = (*arr)[small], (*arr)[r]
		QuickSort(arr, l, small-1)
		QuickSort(arr, small+1, r)
	}
}
