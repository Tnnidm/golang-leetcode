package sumOddLengthSubarrays

func sumOddLengthSubarrays(arr []int) int {
	arrLen := len(arr)
	result := 0
	for subLen := 1; subLen <= arrLen; subLen = subLen + 2 {
		for i := 0; i <= arrLen-subLen; i++ {
			for j := 0; j < subLen; j++ {
				result = result + arr[i+j]
			}
		}
	}
	return result
}
