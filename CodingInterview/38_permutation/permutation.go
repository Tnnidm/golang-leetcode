package permutation

func permutation(s string) []string {
	index := 0
	result := []string{}
	dfs([]byte(s), index, &result)
	return result
}

func dfs(s []byte, index int, result *[]string) {
	if index == len(s)-1 {
		*result = append(*result, string(s))
	} else {
		indexMap := map[byte]int{}
		for i := index + 1; i < len(s); i++ {
			indexMap[s[i]] = i
		}
		dfs(s, index+1, result)
		for _, i := range indexMap {
			if s[index] != s[i] {
				s[i], s[index] = s[index], s[i]
				dfs(s, index+1, result)
				s[i], s[index] = s[index], s[i]
			}
		}
	}
}
