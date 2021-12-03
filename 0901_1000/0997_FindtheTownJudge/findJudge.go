package findJudge

func findJudge(n int, trust [][]int) int {
	trustInfo := make([]int, n)
	for i := 0; i < len(trust); i++ {
		// means trust[i][0] trust trust[i][1]
		trustInfo[trust[i][0]-1]--
		trustInfo[trust[i][1]-1]++
	}
	// fmt.Println(trustInfo)
	for i := 0; i < n; i++ {
		if trustInfo[i] == n-1 {
			return i + 1
		}
	}
	return -1
}
