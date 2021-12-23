package maxPower

func maxPower(s string) int {
	energy := 1
	continuity := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continuity++
			if continuity > energy {
				energy = continuity
			}
		} else {
			continuity = 1
		}
	}
	return energy
}
