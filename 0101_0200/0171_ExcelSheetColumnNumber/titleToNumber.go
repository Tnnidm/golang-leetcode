package titleToNumber

func titleToNumber(columnTitle string) int {
	result := 0
	for _, abyte := range columnTitle {
		result = 26*result + int(abyte-'A') + 1
	}
	return result
}
