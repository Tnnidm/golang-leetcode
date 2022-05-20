package calcEquation

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	variable2Index := map[string]int{}
	variableNum := 0
	for _, equation := range equations {
		if _, ok := variable2Index[equation[0]]; !ok {
			variable2Index[equation[0]] = variableNum
			variableNum++
		}
		if _, ok := variable2Index[equation[1]]; !ok {
			variable2Index[equation[1]] = variableNum
			variableNum++
		}
	}
	calculationGraph := make([][]float64, variableNum)
	for i := 0; i < variableNum; i++ {
		calculationGraph[i] = make([]float64, variableNum)
		calculationGraph[i][i] = 1.0
	}
	for i := 0; i < len(equations); i++ {
		x := variable2Index[equations[i][0]]
		y := variable2Index[equations[i][1]]
		calculationGraph[x][y] = values[i]
		calculationGraph[y][x] = 1 / values[i]
	}
	var dfs func(int, int) float64
	var seen []bool
	dfs = func(sourceIndex, targetIndex int) float64 {
		if calculationGraph[sourceIndex][targetIndex] > 0 {
			return calculationGraph[sourceIndex][targetIndex]
		}
		for i := 0; i < variableNum; i++ {
			if i != sourceIndex && !seen[i] && calculationGraph[sourceIndex][i] > 0 {
				seen[i] = true
				temp := dfs(i, targetIndex)
				seen[i] = false
				if temp != -1.0 {
					return calculationGraph[sourceIndex][i] * temp
				}
			}
		}
		return -1.0
	}
	queriesNum := len(queries)
	result := make([]float64, queriesNum)
	for i := 0; i < queriesNum; i++ {
		sourceIndex, sourceExist := variable2Index[queries[i][0]]
		targetIndex, targetExist := variable2Index[queries[i][1]]
		if !sourceExist || !targetExist {
			result[i] = -1.0
		} else {
			seen = make([]bool, variableNum)
			result[i] = dfs(sourceIndex, targetIndex)
		}
	}
	return result
}
