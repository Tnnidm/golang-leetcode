package validPath

func validPath(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}
	linkTable := make(map[int][]int, n)
	for i := 0; i < len(edges); i++ {
		linkTable[edges[i][0]] = append(linkTable[edges[i][0]], edges[i][1])
		linkTable[edges[i][1]] = append(linkTable[edges[i][1]], edges[i][0])
	}
	queue := []int{start}
	exist := make(map[int]struct{}, n)
	for len(queue) != 0 {
		for _, neigh := range linkTable[queue[0]] {
			if neigh == end {
				return true
			}
			if _, ok := exist[queue[0]]; !ok {
				queue = append(queue, neigh)
			}
		}
		exist[queue[0]] = struct{}{}
		queue = queue[1:]
	}
	return false
}
