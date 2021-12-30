package containsPattern

func containsPattern(arr []int, m int, k int) bool {
	arrLen := len(arr)
	if arrLen < m*k {
		return false
	}
	var found bool
	for begin := 0; begin <= arrLen-m*k; begin++ {
		found = true
		for i := 1; i < k && found; i++ {
			for j := 0; j < m; j++ {
				if arr[begin+m*i+j] != arr[begin+m*(i-1)+j] {
					found = false
					break
				}
			}
		}
		if found {
			return true
		}
	}
	return false
}
