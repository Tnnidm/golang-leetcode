package diStringMatch

func diStringMatch(s string) []int {
	N := len(s)
	begin, end := 0, N
	result := make([]int, N+1)
	for i := 0; i < N; i++ {
		if s[i] == 'I' {
			result[i] = begin
			begin++
		} else {
			result[i] = end
			end--
		}
	}
	result[N] = begin
	return result
}
