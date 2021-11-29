package binaryGap

func binaryGap(n int) int {
	maxDis := 0
	lastOneExist := false
	lastOneIndex := 0
	for i := 0; n != 0; i++ {
		if n%2 == 1 {
			if lastOneExist {
				maxDis = max(maxDis, i-lastOneIndex)
			}
			lastOneExist = true
			lastOneIndex = i
		}
		n = n / 2
	}
	return maxDis
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
