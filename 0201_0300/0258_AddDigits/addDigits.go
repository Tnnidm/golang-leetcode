package addDigits

func addDigits(num int) int {
	newnum := 0
	for num >= 10 {
		newnum = 0
		remainer := 0
		for num != 0 || remainer != 0 {
			remainer = num % 10
			newnum += remainer
			num = num / 10
		}
		num = newnum
	}
	return num
}
