package minOperations

func minOperations(s string) int {
	single := 0
	double := 0
	for i, v := range s {
		if v == '0' {
			if i%2 == 0 {
				single++
			} else {
				double++
			}
		} else {
			if i%2 == 0 {
				double++
			} else {
				single++
			}
		}
	}
	if single <= double {
		return single
	}
	return double
}
