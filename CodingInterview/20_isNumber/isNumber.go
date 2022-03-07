package isNumber

func isNumber(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}
	begin, end := 0, sLen-1
	for s[begin] == ' ' && begin < end {
		begin++
	}
	for s[end] == ' ' && begin < end {
		end--
	}
	s = s[begin : end+1]
	sLen = len(s)
	if sLen == 0 {
		return false
	}
	index := 0
	numeric := isInt(s, &index)
	if index < sLen && s[index] == '.' {
		index++
		numeric = isUnsignedInt(s, &index) || numeric
	}
	if index < sLen && (s[index] == 'E' || s[index] == 'e') {
		index++
		if index == sLen {
			return false
		}
		numeric = isInt(s, &index) && numeric
	}
	return numeric && index == sLen
}

func isUnsignedInt(s string, index *int) bool {
	flag := false
	for *index < len(s) && s[*index] >= '0' && s[*index] <= '9' {
		*index++
		flag = true
	}
	return flag
}

func isInt(s string, index *int) bool {
	if s[*index] == '+' || s[*index] == '-' {
		*index++
	}
	return isUnsignedInt(s, index)
}
