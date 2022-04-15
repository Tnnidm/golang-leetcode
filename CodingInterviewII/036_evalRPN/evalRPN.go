package evalRPN

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	stackLen := 0
	for _, token := range tokens {
		switch token {
		case "+":
			stack[stackLen-2] += stack[stackLen-1]
			stack = stack[:stackLen-1]
			stackLen--
		case "-":
			stack[stackLen-2] -= stack[stackLen-1]
			stack = stack[:stackLen-1]
			stackLen--
		case "*":
			stack[stackLen-2] *= stack[stackLen-1]
			stack = stack[:stackLen-1]
			stackLen--
		case "/":
			stack[stackLen-2] /= stack[stackLen-1]
			stack = stack[:stackLen-1]
			stackLen--
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
			stackLen++
		}
	}
	return stack[0]
}
