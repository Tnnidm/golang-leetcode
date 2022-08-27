package matchRegularExpression

// func isMatch(s string, p string) bool {
// 	sLen := len(s)
// 	pLen := len(p)
// 	if sLen == 0 {
// 		if pLen == 0 {
// 			return true
// 		} else if pLen >= 2 && p[1] == '*' {
// 			return isMatch(s, p[2:])
// 		}
// 	}
// 	if sLen == 0 || pLen == 0 {
// 		return false
// 	}
// 	if pLen >= 2 && p[1] == '*' {
// 		if pLen >= 4 && p[0] == p[2] && p[1] == p[3] {
// 			return isMatch(s, p[2:])
// 		}
// 		if p[0] == s[0] || p[0] == '.' {
// 			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
// 		} else {
// 			return isMatch(s, p[2:])
// 		}
// 	}
// 	if p[0] == s[0] || p[0] == '.' {
// 		return isMatch(s[1:], p[1:])
// 	} else {
// 		return false
// 	}
// }

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	match := func(i int, j int) bool {
		if i == 0 {
			return false
		}
		return p[j-1] == '.' || s[i-1] == p[j-1]
	}
	for i := 0; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if match(i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}
		}
	}
	return dp[sLen][pLen]
}
