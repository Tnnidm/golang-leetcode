package combine

func combine(n int, k int) [][]int {
	result := [][]int{}
	combination := make([]int, k)
	var dfs func(index int)
	dfs = func(index int) {
		if index == k {
			result = append(result, append([]int{}, combination...))
			return
		}
		var start int
		if index == 0 {
			start = 1
		} else {
			start = combination[index-1] + 1
		}
		// fmt.Println(start, n-k+index+1)
		for i := start; i <= n-k+index+1; i++ {
			combination[index] = i
			dfs(index + 1)
		}
	}
	dfs(0)
	return result
}
