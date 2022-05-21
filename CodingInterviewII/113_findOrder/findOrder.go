package findOrder

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses) //邻接表
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	indegree := make([]int, numCourses) //入度表
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		indegree[prerequisites[i][0]]++
	}

	var queue []int
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		result = append(result, cur)
		for _, nxt := range graph[cur] {
			indegree[nxt]--
			if indegree[nxt] == 0 {
				queue = append(queue, nxt)
			}
		}
	}
	//如果访问数量不等于节点数量，则证明成环
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
