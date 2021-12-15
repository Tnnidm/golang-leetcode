package uniqueOccurrences

func uniqueOccurrences(arr []int) bool {
	statmap := make(map[int]int)
	for _, v := range arr {
		statmap[v]++
	}
	nummap := make(map[int]struct{})
	for _, v := range statmap {
		if _, ok := nummap[v]; ok {
			return false
		}
		nummap[v] = struct{}{}
	}
	return true
}
