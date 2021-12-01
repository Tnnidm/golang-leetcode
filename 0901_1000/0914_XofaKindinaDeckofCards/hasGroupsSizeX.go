package hasGroupsSizeX

func hasGroupsSizeX(deck []int) bool {
	counter := make(map[int]int)
	for _, v := range deck {
		counter[v]++
	}
	minv := 10000
	for _, v := range counter {
		if minv > v {
			minv = v
		}
	}
	if minv < 2 {
		return false
	}
	flag := true
	for i := 2; i <= minv; i++ {
		flag = true
		for _, v := range counter {
			if v%i != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}
