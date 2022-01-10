package minimumMoves

func minimumMoves(s string) int {
	result := 0
	for i := 0; i < len(s); {
		if s[i] == 'X' {
			result++
			i += 3
		} else {
			i++
		}
	}
	return result
}
