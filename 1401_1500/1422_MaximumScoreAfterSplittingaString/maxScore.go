package maxScore

func maxScore(s string) int {
	sLen := len(s)
	var score, max int
	for split := 1; split < sLen; split++ {
		score = 0
		for i := 0; i < split; i++ {
			if s[i] == '0' {
				score++
			}
		}
		for i := split; i < sLen; i++ {
			if s[i] == '1' {
				score++
			}
		}
		if score > max {
			max = score
		}
	}
	return max
}
