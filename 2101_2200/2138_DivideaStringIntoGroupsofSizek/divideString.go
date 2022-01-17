package divideString

func divideString(s string, k int, fill byte) []string {
	sLen := len(s)
	var result []string
	if sLen%k == 0 {
		result = make([]string, sLen/k)
		for i := 0; i < sLen/k; i++ {
			result[i] = s[k*i : k*i+k]
		}
	} else {
		result = make([]string, sLen/k+1)
		for i := 0; i < sLen/k; i++ {
			result[i] = s[k*i : k*i+k]
		}
		lastOne := make([]byte, k)
		for i := 0; i < sLen%k; i++ {
			lastOne[i] = s[sLen-sLen%k+i]
		}
		for i := sLen % k; i < k; i++ {
			lastOne[i] = fill
		}
		result[sLen/k] = string(lastOne)
	}
	return result
}
