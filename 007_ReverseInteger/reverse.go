package reverse

// // Solution 1 调用math包
// import "math"
// func reverse(x int) int {
// 	flag := false
// 	var newnum int
// 	if x < 0 {
// 		x = -x
// 		flag = true
// 	}
// 	for {
// 		newnum = newnum*10 + x%10

// 		x = x / 10
// 		if x == 0 {
// 			break
// 		}
// 	}
// 	if flag {
// 		newnum = -newnum
// 	}
// 	if newnum < math.MinInt32 || newnum > math.MaxInt32 {
// 		return 0
// 	}
// 	return newnum
// }

// Solution 2 不调用math包，内存小上一种4kb
func reverse(x int) int {
	flag := false
	var newnum int
	var INT_MAX = int(^uint32(0) >> 1)
	var INT_MIN = ^int(^uint32(0) >> 1)
	if x < 0 {
		x = -x
		flag = true
	}
	for {
		newnum = newnum*10 + x%10

		x = x / 10
		if x == 0 {
			break
		}
	}
	if flag {
		newnum = -newnum
	}
	if newnum < INT_MIN || newnum > INT_MAX {
		return 0
	}
	return newnum
}
