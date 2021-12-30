package modifyString

func modifyString(s string) string {
	sLen := len(s)
	sByte := []byte(s)
	ques := byte('?')
	for i := 0; i < sLen; i++ {
		if sByte[i] == ques {
			replace := byte('a')
			for (i > 0 && sByte[i-1] == replace) || (i < sLen-1 && sByte[i+1] == replace) {
				replace += 1
			}
			sByte[i] = replace
		}
	}
	return string(sByte)
}
