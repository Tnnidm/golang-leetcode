package countStudents

func countStudents(students []int, sandwiches []int) int {
	oneCount, zeroCount := 0, 0
	N := len(students)
	for i := 0; i < N; i++ {
		if students[i] == 1 {
			oneCount++
		} else {
			zeroCount++
		}
	}
	for i := 0; i < N; i++ {
		if sandwiches[i] == 1 {
			if oneCount > 0 {
				oneCount--
			} else {
				return N - i
			}
		} else {
			if zeroCount > 0 {
				zeroCount--
			} else {
				return N - i
			}
		}
	}
	return 0
}
