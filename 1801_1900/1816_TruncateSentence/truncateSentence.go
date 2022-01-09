package truncateSentence

func truncateSentence(s string, k int) string {
	for i := range s {
		if s[i] == ' ' {
			k--
			if k == 0 {
				s = s[:i]
				break
			}
		}
	}
	return s
}
