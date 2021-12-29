package makeGood

func makeGood(s string) string {
	exit := true
	sByte := []byte(s)
	diff := byte('a' - 'A')
	for exit {
		exit = false
		for i := 0; i < len(sByte)-1; i++ {
			if sByte[i+1]-sByte[i] == diff || sByte[i]-sByte[i+1] == diff {
				sByte = append(sByte[:i], sByte[i+2:]...)
				exit = true
				break
			}
		}
	}
	return string(sByte)
}
