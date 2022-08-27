package permutation

func permutation(s string) []string {
	result := []string{}
	permut := []byte(s)
	sLen := len(s)
	var dfs func(int)
	dfs = func(index int) {
		if index == sLen-1 {
			result = append(result, string(permut))
			return
		}
		dfs(index + 1)
		swap := map[byte]int{}
		for i := index + 1; i < sLen; i++ {
			swap[permut[i]] = i
		}
		for _, i := range swap {
			if permut[index] != permut[i] {
				permut[i], permut[index] = permut[index], permut[i]
				dfs(index + 1)
				permut[i], permut[index] = permut[index], permut[i]
			}
		}
	}
	dfs(0)
	return result
}
