package fairCandySwap

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceSum, aliceNum := 0, len(aliceSizes)
	bobSum, bobNum := 0, len(bobSizes)
	var i int
	for i = 0; i < aliceNum; i++ {
		aliceSum += aliceSizes[i]
	}
	for i = 0; i < bobNum; i++ {
		bobSum += bobSizes[i]
	}
	diff := (bobSum - aliceSum) / 2
	b2aMap := make(map[int]struct{}, bobNum)
	for i = 0; i < bobNum; i++ {
		b2aMap[bobSizes[i]-diff] = struct{}{}
	}
	for i = 0; i < aliceNum; i++ {
		if _, ok := b2aMap[aliceSizes[i]]; ok {
			return []int{aliceSizes[i], aliceSizes[i] + diff}
		}
	}
	return []int{aliceSizes[i], aliceSizes[i] + diff}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
