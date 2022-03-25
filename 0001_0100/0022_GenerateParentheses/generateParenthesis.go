package generateParenthesis

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	result := []string{}
	str := make([]byte, 2*n)
	dfs(n, 0, 0, &str, &result)
	return result
}

func dfs(left, right, strIndex int, str *[]byte, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, string(*str))
		return
	}
	if left > 0 {
		(*str)[strIndex] = '('
		dfs(left-1, right+1, strIndex+1, str, result)
	}
	if right > 0 {
		(*str)[strIndex] = ')'
		dfs(left, right-1, strIndex+1, str, result)
	}
}
