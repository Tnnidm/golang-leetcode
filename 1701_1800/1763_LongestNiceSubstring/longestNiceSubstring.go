package longestNiceSubstring

type NiceChar struct {
	upper bool
	lower bool
}

func longestNiceSubstring(s string) string {
	sLen := len(s)
	for length := sLen; length >= 1; length-- {
		for i := 0; i <= sLen-length; i++ {
			if isNice(s[i : i+length]) {
				return s[i : i+length]
			}
		}
	}
	return ""
}

func isNice(subString string) bool {
	nice := make([]NiceChar, 26)
	for _, v := range subString {
		if v >= 'a' && v <= 'z' {
			nice[v-'a'].lower = true
		}
		if v >= 'A' && v <= 'Z' {
			nice[v-'A'].upper = true
		}
	}
	for i := 0; i < 26; i++ {
		if nice[i].upper != nice[i].lower {
			return false
		}
	}
	return true
}
