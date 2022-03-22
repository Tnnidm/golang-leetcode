package winnerOfGame

func winnerOfGame(colors string) bool {
	count := 0
	for i := 1; i < len(colors)-1; i++ {
		if colors[i] == 'A' && colors[i-1] == 'A' && colors[i+1] == 'A' {
			count++
		} else if colors[i] == 'B' && colors[i-1] == 'B' && colors[i+1] == 'B' {
			count--
		}
	}
	return count > 0
}
