package reverseString

func reverseString(s []byte) {
	for pos, neg := 0, len(s)-1; pos < neg; {
		s[pos], s[neg] = s[neg], s[pos]
		pos++
		neg--
	}
}
