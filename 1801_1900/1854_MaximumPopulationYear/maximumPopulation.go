package maximumPopulation

func maximumPopulation(logs [][]int) int {
	populationChange := make([]int, 101)
	for i := range logs {
		populationChange[logs[i][0]-1950]++
		populationChange[logs[i][1]-1950]--
	}
	maxPopu := populationChange[0]
	popu := populationChange[0]
	maxPopuYear := 0
	for i := 1; i < 101; i++ {
		popu = popu + populationChange[i]
		if popu > maxPopu {
			maxPopu = popu
			maxPopuYear = i
		}
	}
	return maxPopuYear + 1950
}
