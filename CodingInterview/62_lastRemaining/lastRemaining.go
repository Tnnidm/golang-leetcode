package lastRemaining

// // 方法1: 模拟
// func lastRemaining(n int, m int) int {
// 	assist := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		assist[i] = i
// 	}
// 	index := 0
// 	for len(assist) != 1 {
// 		index += (m - 1) % n
// 		index %= n
// 		assist = append(assist[:index], assist[index+1:]...)
// 		n--
// 	}
// 	return assist[0]
// }

// 方法2: 数学
func lastRemaining(n int, m int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
