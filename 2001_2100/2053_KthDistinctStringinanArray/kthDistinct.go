package kthDistinct

func kthDistinct(a []string, k int) string {
	stringCount := make(map[string]int, len(a))
	for _, str := range a {
		stringCount[str]++
	}
	for _, str := range a {
		if stringCount[str] == 1 {
			k--
			if k == 0 {
				return str
			}
		}
	}
	return ""
}
