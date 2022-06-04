package canFinish

// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegree[edge[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range graph[cur] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	for _, inde := range inDegree {
		if inde != 0 {
			return false
		}
	}
	return true
}
