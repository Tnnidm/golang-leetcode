package countSegments

func countSegments(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			result++
		}
	}
	return result
}
