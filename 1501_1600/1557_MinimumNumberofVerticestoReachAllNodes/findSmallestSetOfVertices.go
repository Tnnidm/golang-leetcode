package findSmallestSetOfVertices

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	counting := make([]bool, n)
	for i := 0; i < len(edges); i++ {
		counting[edges[i][1]] = true
	}
	result := []int{}
	for i := 0; i < n; i++ {
		if !counting[i] {
			result = append(result, i)
		}
	}
	return result
}
