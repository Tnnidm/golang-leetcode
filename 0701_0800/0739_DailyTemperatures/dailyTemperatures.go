package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	temperatureLen := len(temperatures)
	ans := make([]int, temperatureLen)
	stack := []int{}
	stackLen := 0
	for i := 0; i < temperatureLen; i++ {
		for stackLen > 0 && temperatures[stack[stackLen-1]] < temperatures[i] {
			ans[stack[stackLen-1]] = i - stack[stackLen-1]
			stackLen--
			stack = stack[:stackLen]
		}
		stack = append(stack, i)
		stackLen++
	}
	return ans
}
