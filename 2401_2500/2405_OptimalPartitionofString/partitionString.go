package partitionString

func partitionString(s string) int {
	count := make([]bool, 26)
	result := 1
	for i := range s {
		if count[s[i]-'a'] {
			result++
			count = make([]bool, 26)
		}
		count[s[i]-'a'] = true
	}
	return result
}
