package largeGroupPositions

func largeGroupPositions(s string) [][]int {
	sLen := len(s)
	if sLen < 3 {
		return [][]int{}
	}
	result := [][]int{}
	begin := 0
	for i := 0; i < sLen; i++ {
		if i == sLen-1 || s[i] != s[i+1] {
			if i-begin+1 >= 3 {
				result = append(result, []int{begin, i})
			}
			begin = i + 1
		}
	}
	return result
}
