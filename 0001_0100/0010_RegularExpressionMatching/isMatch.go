package isMatch

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	if sLen == 0 {
		if pLen == 0 {
			return true
		} else if pLen >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		}
	}
	if sLen == 0 || pLen == 0 {
		return false
	}
	if pLen >= 2 && p[1] == '*' {
		if pLen >= 4 && p[0] == p[2] && p[1] == p[3] {
			return isMatch(s, p[2:])
		}
		if p[0] == s[0] || p[0] == '.' {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	}
	if p[0] == s[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}
