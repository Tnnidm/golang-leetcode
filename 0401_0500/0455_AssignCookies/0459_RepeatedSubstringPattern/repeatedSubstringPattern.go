package repeatedSubstringPattern

func repeatedSubstringPattern(s string) bool {
	sLen := len(s)
	for sessionLen := 1; sessionLen <= sLen/2; sessionLen++ {
		if sLen%sessionLen == 0 {
			breakFlag := false
			for i := 0; i < sLen/sessionLen; i++ {
				for j := 0; j < sessionLen; j++ {
					if s[i*sessionLen+j] != s[j] {
						breakFlag = true
						break
					}
				}
				if breakFlag {
					break
				}
			}
			if breakFlag == false {
				return true
			}
		}
	}
	return false
}
