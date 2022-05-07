package partition

func partition(s string) [][]string {
	sLen := len(s)
	start2end := make([][]int, sLen)
	// odd
	for i := 1; i < sLen-1; i++ {
		for j := 1; i-j >= 0 && i+j < sLen && s[i-j] == s[i+j]; j++ {
			start2end[i-j] = append(start2end[i-j], i+j)
		}
	}
	// even
	for i := 0; i < sLen-1; i++ {
		for j := 0; i-j >= 0 && i+1+j < sLen && s[i-j] == s[i+1+j]; j++ {
			start2end[i-j] = append(start2end[i-j], i+1+j)
		}
	}
	// backtrack
	result := [][]string{}
	combination := []string{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == sLen {
			result = append(result, append([]string{}, combination...))
			return
		}
		combinationLen := len(combination)
		// 只包含一个字符的情况
		combination = append(combination, s[index:index+1])
		dfs(index + 1)
		combination = combination[:combinationLen]
		// 长度超过1个字符的情况
		for _, end := range start2end[index] {
			combination = append(combination, s[index:end+1])
			dfs(end + 1)
			combination = combination[:combinationLen]
		}
	}
	dfs(0)
	return result
}
