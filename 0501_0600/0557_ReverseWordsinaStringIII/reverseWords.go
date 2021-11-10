package reverseWords

func reverseWords(s string) string {
	sLen := len(s)
	if sLen == 1 {
		return s
	} else {
		var begin int
		sByte := []byte(s)
		for i := 0; i < sLen; i++ {
			if i == 0 || (sByte[i-1] == ' ' && sByte[i] != ' ') {
				begin = i
			}
			if i == sLen-1 || (sByte[i] != ' ' && sByte[i+1] == ' ') {
				for j := begin; j <= (begin+i)/2; j++ {
					sByte[j], sByte[i+begin-j] = sByte[i+begin-j], sByte[j]
				}
			}
		}
		return string(sByte)
	}
}
