package duplicateZeros

func duplicateZeros(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		if arr[i] == 0 {
			for j := arrLen - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
}
