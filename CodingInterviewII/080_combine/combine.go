package combine

func combine(n int, k int) [][]int {
	result := [][]int{}
	temp := make([]int, 0, k)
	tempLen := 0
	var dfs func(int)
	dfs = func(number int) {
		if tempLen == k {
			result = append(result, append(make([]int, 0, k), temp...))
			return
		}
		if number > n {
			return
		}
		dfs(number + 1)
		temp = append(temp, number)
		tempLen++
		dfs(number + 1)
		tempLen--
		temp = temp[:tempLen]
	}
	dfs(1)
	return result
}
