package minWindow

func minWindow(s string, t string) string {
	var sCount, tCount [58]int
	for i := range t {
		tCount[t[i]-'A']++
	}
	left, right := 0, 0
	result := ""
	resultLen := len(s) + 1
	for {
		if !cover(sCount, tCount) {
			if right < len(s) {
				sCount[s[right]-'A']++
				right++
			} else {
				break
			}
		} else {
			if right-left < resultLen {
				resultLen = right - left + 1
				result = s[left:right]
			}
			sCount[s[left]-'A']--
			left++
		}
	}
	return result
}

func cover(sCount, tCount [58]int) bool {
	for i := range sCount {
		if sCount[i] < tCount[i] {
			return false
		}
	}
	return true
}
