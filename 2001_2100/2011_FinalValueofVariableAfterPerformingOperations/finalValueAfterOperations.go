package finalValueAfterOperations

func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, operation := range operations {
		if operation[0] == '+' || operation[2] == '+' {
			result++
		} else {
			result--
		}
	}
	return result
}
