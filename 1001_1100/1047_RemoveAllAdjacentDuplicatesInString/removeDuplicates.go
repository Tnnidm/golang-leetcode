package removeDuplicates

func removeDuplicates(s string) string {
	sByte := []byte{}
	sLen := 0
	for i := 0; i < len(s); i++ {
		if sLen == 0 || sByte[sLen-1] != s[i] {
			sByte = append(sByte, s[i])
			sLen++
		} else {
			sByte = sByte[:sLen-1]
			sLen--
		}
	}
	return string(sByte)
}
