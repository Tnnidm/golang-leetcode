package findSpecialInteger

func findSpecialInteger(arr []int) int {
	arrLen := len(arr)
	threshold := arrLen / 4
	if arrLen < 4 {
		return arr[0]
	}
	counter := 1
	for i := 1; i < arrLen; i++ {
		if arr[i] == arr[i-1] {
			counter++
			if counter > threshold {
				return arr[i]
			}
		} else {
			counter = 1
		}
	}
	return arr[0]
}
