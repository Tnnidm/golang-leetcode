package detectCapitalUse

func detectCapitalUse(word string) bool {
	wordLen := len(word)
	if wordLen == 1 {
		return true
	} else {
		if word[1] >= 'a' && word[1] <= 'z' {
			for i := 2; i < wordLen; i++ {
				if word[i] >= 'A' && word[i] <= 'Z' {
					return false
				}
			}
			return true
		} else if word[1] >= 'A' && word[1] <= 'Z' {
			for i := 0; i < wordLen; i++ {
				if word[i] >= 'a' && word[i] <= 'z' {
					return false
				}
			}
			return true
		}
	}
	return false
}
