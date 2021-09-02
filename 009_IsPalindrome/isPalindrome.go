package isPalindrome

// method 1, using the same method like 007 reverse integer
// defeat 88.13% users in time and 74.04% users in memory
func isPalindrome(x int) bool {
	var result bool
	if x < 0 {
		result = false
	} else {
		var newnum int
		var oldnum = x
		for {
			newnum = newnum*10 + oldnum%10
			oldnum = oldnum / 10
			if oldnum == 0 {
				break
			}
		}
		if newnum == x {
			result = true
		} else {
			result = false
		}
	}
	return result
}
