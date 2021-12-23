package buildArray

func buildArray(target []int, n int) []string {
	ptr := 0
	result := []string{}
	for i := 1; i <= n; i++ {
		if i == target[ptr] {
			result = append(result, "Push")
			if i == target[len(target)-1] {
				break
			}
			ptr++
		} else {
			result = append(result, "Push")
			result = append(result, "Pop")
		}
	}
	return result
}
