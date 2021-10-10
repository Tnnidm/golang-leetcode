package convertToTitle

func convertToTitle(columnNumber int) string {
	bytes := []byte{}
	remain := (columnNumber-1)%26 + 1
	columnNumber = (columnNumber - 1) / 26
	for remain != 0 || columnNumber != 0 {
		bytes = append([]byte{uint8(remain + 64)}, bytes...)
		remain = (columnNumber-1)%26 + 1
		columnNumber = (columnNumber - 1) / 26
	}
	return string(bytes)
}
