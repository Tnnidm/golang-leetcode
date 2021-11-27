package main

func shortestToChar(s string, c byte) []int {
	lastcIndex := -10000
	sLen := len(s)
	result := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		if s[i] == c {
			lastcIndex = i
		}
		result[i] = i - lastcIndex
	}
	lastcIndex = 10000
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == c {
			lastcIndex = i
		}
		result[i] = min(result[i], lastcIndex-i)
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
