package restoreString

func restoreString(s string, indices []int) string {
	byteS := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		index := indices[i]
		byteS[index] = s[i]
	}
	return string(byteS)
}
