package isOneBitCharacter

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if i == len(bits)-1 && bits[i] == 0 {
			return true
		} else {
			if bits[i] == 1 {
				i++
			}
		}
	}
	return false
}
