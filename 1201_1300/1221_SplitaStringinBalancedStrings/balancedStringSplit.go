package balancedStringSplit

func balancedStringSplit(s string) int {
	LRNum := 0
	ans := 0
	for _, v := range s {
		if v == 'L' {
			LRNum++
		} else {
			LRNum--
		}
		if LRNum == 0 {
			ans++
		}
	}
	return ans
}
