package maxDepth

func maxDepth(s string) int {
	depth := 0
	maxDepth := 0
	for _, v := range s {
		if v == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if v == ')' {
			depth--
		}
	}
	return maxDepth
}
