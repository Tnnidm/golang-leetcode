package countDigitOne

func countDigitOne(n int) int {
	ans := 0
	thisDigit := 1
	for n >= thisDigit {
		nextDigit := 10 * thisDigit
		ans += (n/(nextDigit))*thisDigit + min(max(n%nextDigit-thisDigit+1, 0), thisDigit)
		thisDigit = nextDigit
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
