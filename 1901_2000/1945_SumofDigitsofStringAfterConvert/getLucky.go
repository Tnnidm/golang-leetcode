package getLucky

func getLucky(s string, k int) int {
	num := 0
	for _, v := range s {
		switch v {
		case 'a', 'j':
			num += 1
		case 'b', 'k', 't':
			num += 2
		case 'c', 'l', 'u':
			num += 3
		case 'd', 'm', 'v':
			num += 4
		case 'e', 'n', 'w':
			num += 5
		case 'f', 'o', 'x':
			num += 6
		case 'g', 'p', 'y':
			num += 7
		case 'h', 'q', 'z':
			num += 8
		case 'i', 'r':
			num += 9
		case 's':
			num += 10
		}
	}
	k--
	nextNum := 0
	for ; k != 0; k-- {
		nextNum = 0
		for num != 0 {
			nextNum += num % 10
			num = num / 10
		}
		num = nextNum
	}
	return num
}
