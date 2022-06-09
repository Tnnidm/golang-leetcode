package calcEquation

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	variableSet := map[string]int{}
	count := 0
	for _, equation := range equations {
		if _, ok := variableSet[equation[0]]; !ok {
			variableSet[equation[0]] = count
			count++
		}
		if _, ok := variableSet[equation[1]]; !ok {
			variableSet[equation[1]] = count
			count++
		}
	}
	graph := make([][]float64, count)
	for i := 0; i < count; i++ {
		graph[i] = make([]float64, count)
		graph[i][i] = 1.0
	}
	for index, equation := range equations {
		graph[variableSet[equation[0]]][variableSet[equation[1]]] = values[index]
		graph[variableSet[equation[1]]][variableSet[equation[0]]] = 1 / values[index]
	}
	// Floyd算法
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] != 0 && graph[k][j] != 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}
	result := make([]float64, len(queries))
	for resultIndex, query := range queries {
		index1, ok1 := variableSet[query[0]]
		index2, ok2 := variableSet[query[1]]
		if !(ok1 && ok2) || graph[index1][index2] == 0 {
			result[resultIndex] = -1.0
		} else {
			result[resultIndex] = graph[index1][index2]
		}
	}
	return result
}
