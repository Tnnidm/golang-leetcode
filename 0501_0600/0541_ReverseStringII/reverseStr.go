package reverseStr

func reverseStr(s string, k int) string {
	sLen := len(s)
	sByte := []byte(s)
	for i := 0; i <= sLen/(2*k); i++ {
		subLen := min(k, sLen-i*2*k)
		for j := 0; j < subLen/2; j++ {
			sByte[2*k*i+j], sByte[2*k*i+subLen-1-j] = sByte[2*k*i+subLen-1-j], sByte[2*k*i+j]
		}
	}
	return string(sByte)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
