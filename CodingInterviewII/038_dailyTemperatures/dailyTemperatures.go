package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	temperaturesLen := len(temperatures)
	stack := make([]int, 0, 64) // 这里的64只是预估的，预估一个值可以让内存更省
	stackLen := 0
	result := make([]int, temperaturesLen)
	for i := 0; i < temperaturesLen; i++ {
		for stackLen > 0 && temperatures[stack[stackLen-1]] < temperatures[i] {
			result[stack[stackLen-1]] = i - stack[stackLen-1]
			stackLen--
			stack = stack[:stackLen]
		}
		stack = append(stack, i)
		stackLen++
	}
	return result
}
