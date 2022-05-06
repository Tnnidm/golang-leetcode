package generateParenthesis

func generateParenthesis(n int) []string {
	result := []string{}
	combanation := make([]byte, 2*n)
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left == 0 && right == 0 {
			result = append(result, string(combanation))
			return
		}
		index := 2*n - 2*left - right
		if left > 0 {
			combanation[index] = '('
			dfs(left-1, right+1)
		}
		if right > 0 {
			combanation[index] = ')'
			dfs(left, right-1)
		}
	}
	dfs(n, 0)
	return result
}
