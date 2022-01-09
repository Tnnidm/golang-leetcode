package replaceDigits

func replaceDigits(s string) string {
	sByte := []byte(s)
	for i := 1; i < len(sByte); i += 2 {
		sByte[i] = sByte[i-1] + sByte[i] - '0'
	}
	return string(sByte)
}
