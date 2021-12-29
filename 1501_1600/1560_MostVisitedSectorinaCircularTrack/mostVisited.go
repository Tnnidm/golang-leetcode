package mostVisited

func mostVisited(n int, rounds []int) []int {
	m := len(rounds)
	start := rounds[0]
	end := rounds[m-1]
	if start <= end {
		result := make([]int, end-start+1)
		for i := start; i <= end; i++ {
			result[i-start] = i
		}
		return result
	} else {
		result := make([]int, end+1+n-start)
		index := 0
		for i := 1; i <= end; i++ {
			result[index] = i
			index++
		}
		for i := start; i <= n; i++ {
			result[index] = i
			index++
		}
		return result
	}
}
