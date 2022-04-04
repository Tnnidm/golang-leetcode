package removeComments

func removeComments(source []string) []string {
	blockCommentFlag := false
	result := []string{}
	beyondComment := ""
	for i := 0; i < len(source); i++ {
		for j := 1; j < len(source[i]); j++ {
			if !blockCommentFlag && source[i][j-1] == '/' && source[i][j] == '/' {
				source[i] = source[i][:j-1]
				break
			}
			if !blockCommentFlag && source[i][j-1] == '/' && source[i][j] == '*' {
				beyondComment = source[i][:j-1]
				blockCommentFlag = true
				j++
				continue
			}
			if blockCommentFlag && source[i][j-1] == '*' && source[i][j] == '/' {
				source[i] = beyondComment + source[i][j+1:]
				j = len(beyondComment)
				beyondComment = ""
				blockCommentFlag = false
			}
		}
		if !blockCommentFlag && source[i] != "" {
			result = append(result, source[i])
		}
	}
	return result
}
