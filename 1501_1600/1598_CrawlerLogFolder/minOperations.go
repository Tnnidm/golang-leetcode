package minOperations

func minOperations(logs []string) int {
	result := 0
	for _, log := range logs {
		if log == "../" {
			if result > 0 {
				result--
			}
			continue
		}
		if log == "./" {
			continue
		}
		result++
	}
	return result
}
