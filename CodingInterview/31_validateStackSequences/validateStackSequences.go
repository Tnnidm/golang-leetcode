package validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	pushIndex, popIndex := 0, 0
	pushStackLen := len(pushed)
	poppedNum := pushStackLen
	for popIndex < poppedNum {
		if pushed[pushIndex] != popped[popIndex] {
			if pushIndex == pushStackLen-1 {
				return false
			} else {
				pushIndex++
			}
		} else {
			popIndex++
			pushed = append(pushed[:pushIndex], pushed[pushIndex+1:]...)
			if pushIndex != 0 {
				pushIndex--
			}
			pushStackLen--
		}
	}
	return true
}
