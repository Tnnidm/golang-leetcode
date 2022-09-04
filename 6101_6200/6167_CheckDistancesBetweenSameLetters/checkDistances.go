package checkDistances

func checkDistances(s string, distance []int) bool {
	seen := make([]bool, 26)
	for i := range s {
		if !seen[s[i]-'a'] {
			seen[s[i]-'a'] = true
			distance[s[i]-'a'] += i + 1
		} else {
			if distance[s[i]-'a'] != i {
				return false
			}
		}
	}
	return true
}
