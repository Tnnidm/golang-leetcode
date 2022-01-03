package maximumWealth

func maximumWealth(accounts [][]int) int {
	maxWealth, wealth := 0, 0
	for i := 0; i < len(accounts); i++ {
		wealth = 0
		for j := 0; j < len(accounts[0]); j++ {
			wealth = wealth + accounts[i][j]
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
