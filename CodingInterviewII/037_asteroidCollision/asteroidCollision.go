package asteroidCollision

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	stackLen := 0
	for _, asteroid := range asteroids {
		vanishFlag := false
		for asteroid < 0 && stackLen > 0 && stack[stackLen-1] > 0 {
			if stack[stackLen-1] == -asteroid {
				stack = stack[:stackLen-1]
				stackLen--
				vanishFlag = true
				break
			} else if stack[stackLen-1] > -asteroid {
				vanishFlag = true
				break
			} else {
				stack = stack[:stackLen-1]
				stackLen--
			}
		}
		if !vanishFlag {
			stack = append(stack, asteroid)
			stackLen++
		}
	}
	return stack
}
