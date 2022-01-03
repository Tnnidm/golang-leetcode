package countPoints

type rod struct {
	redRing   bool
	greenRing bool
	blueRing  bool
}

func countPoints(rings string) int {
	rods := make([]rod, 10)
	for i := 0; i < len(rings)/2; i++ {
		switch rings[2*i] {
		case 'R':
			rods[rings[2*i+1]-'0'].redRing = true
		case 'G':
			rods[rings[2*i+1]-'0'].greenRing = true
		case 'B':
			rods[rings[2*i+1]-'0'].blueRing = true
		}
	}
	result := 0
	for i := 0; i < 10; i++ {
		if rods[i].redRing && rods[i].greenRing && rods[i].blueRing {
			result++
		}
	}
	return result
}
