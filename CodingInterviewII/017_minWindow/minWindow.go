package minWindow

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	if sLen < tLen {
		return ""
	}
	tCount := make([]int, 58)
	for i := 0; i < tLen; i++ {
		tCount[t[i]-'A']++
	}
	sCount := make([]int, 58)
	left, right := 0, 0
	minLen := 100001
	minLeft, minRight := 0, 0
	for {
		for right < sLen && !include(tCount, sCount) {
			sCount[s[right]-'A']++
			right++
		}
		if !include(tCount, sCount) {
			break
		}
		for left < right && include(tCount, sCount) {
			sCount[s[left]-'A']--
			left++
		}
		if right-left+1 < minLen {
			minLen = right - left + 1
			minLeft = left - 1
			minRight = right
		}
	}
	return s[minLeft:minRight]
}

func include(tCount []int, sCount []int) bool {
	for i := 0; i < 58; i++ {
		if tCount[i] > sCount[i] {
			return false
		}
	}
	return true
}
