package lemonadeChange

func lemonadeChange(bills []int) bool {
	change := make([]int, 2)
	for _, v := range bills {
		switch v {
		case 5:
			change[0]++
		case 10:
			change[1]++
			if change[0] > 0 {
				change[0]--
			} else {
				return false
			}
		case 20:
			if change[1] > 0 && change[0] > 0 {
				change[1]--
				change[0]--
			} else if change[0] >= 3 {
				change[0] = change[0] - 3
			} else {
				return false
			}
		}
	}
	return true
}
