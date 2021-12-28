package countGoodTriplets

func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	arrLen := len(arr)
	for i := 0; i < arrLen-2; i++ {
		for j := i + 1; j < arrLen-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < arrLen; k++ {
					if abs(arr[i]-arr[k]) <= c && abs(arr[j]-arr[k]) <= b {
						result++
					}
				}
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
