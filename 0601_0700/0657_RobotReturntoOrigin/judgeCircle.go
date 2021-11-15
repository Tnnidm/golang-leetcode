package judgeCircle

func judgeCircle(moves string) bool {
	LR, UD := 0, 0
	for _, v := range moves {
		switch v {
		case 'L':
			LR--
		case 'R':
			LR++
		case 'U':
			UD++
		case 'D':
			UD--
		}
	}
	return LR == 0 && UD == 0
}
