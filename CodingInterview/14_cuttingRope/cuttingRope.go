package cuttingRope

// 动态规划
func cuttingRope(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	default:
		result := make([]int, n+1)
		result[2] = 2
		result[3] = 3
		result[4] = 4
		for i := 5; i <= n; i++ {
			for j := 2; j <= i/2; j++ {
				product := result[j] * result[i-j]
				if product > result[i] {
					result[i] = product
				}
			}
		}
		return result[n]
	}
}

// //贪婪
// func cuttingRope(n int) int {
// 	switch n {
// 	case 2:
// 		return 1
// 	case 3:
// 		return 2
// 	case 4:
// 		return 4
// 	default:
// 		result := 1
// 		for ; n > 4; n -= 3 {
// 			result *= 3
// 		}
// 		result *= n
// 		return result
// 	}
// }
