package timeRequiredToBuy

func timeRequiredToBuy(tickets []int, k int) int {
	peoNum := len(tickets)
	time := 0
	for i := 0; ; i++ {
		if tickets[i%peoNum] > 0 {
			tickets[i%peoNum]--
			time++
			if tickets[k] == 0 {
				return time
			}
		}
	}
}
