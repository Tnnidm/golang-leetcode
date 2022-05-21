package sequenceReconstruction

func sequenceReconstruction(org []int, seqs [][]int) bool {
	// 注意判断seqs中元素的唯一性
	// 注意判断org里面是否都覆盖到了
	number := len(org)
	graph := make([][]int, number)
	inDegree := make([]int, number)
	visit := make([]bool, number)
	for _, seq := range seqs {
		exist := map[int]bool{}
		if seq[0] < 1 || seq[0] > number {
			return false
		}
		exist[seq[0]-1] = true
		visit[seq[0]-1] = true
		for i := 1; i < len(seq); i++ {
			if seq[i] < 1 || seq[i] > number {
				return false
			}
			if exist[seq[i]-1] {
				return false
			}
			exist[seq[i]-1] = true
			visit[seq[i]-1] = true
			graph[seq[i-1]-1] = append(graph[seq[i-1]-1], seq[i]-1)
			inDegree[seq[i]-1]++
		}
	}
	queue := []int{}
	for i := 0; i < number; i++ {
		if !visit[i] {
			return false
		}
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	index := 0
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}
		node := queue[0]
		queue = queue[1:]
		if org[index]-1 != node {
			return false
		}
		index++
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return index == number
}
