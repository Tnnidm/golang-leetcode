package distributeCandies

func distributeCandies(candyType []int) int {
	candyNum := len(candyType)
	typeMap := make(map[int]struct{})
	typeNum := 0
	for i := 0; i < candyNum; i++ {
		if _, ok := typeMap[candyType[i]]; !ok {
			typeMap[candyType[i]] = struct{}{}
			typeNum++
		}
	}
	candyNum = candyNum / 2
	if typeNum > candyNum {
		return candyNum
	} else {
		return typeNum
	}
}
