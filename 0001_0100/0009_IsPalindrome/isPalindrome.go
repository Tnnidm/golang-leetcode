package isPalindrome

// my method , using the same method like 007 reverse integer
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

// // 优化时间复杂度
// // learned from others' code, convert int into stirng
// // having lower time complexity but higher space complexity
// // defeat 97% users in time and 32% users in memory
// import "strconv"
// func isPalindrome(x int) bool {
//     str := strconv.Itoa(x)
//     strLen := len(str)
//     head, tail := 0, strLen-1
//     result := true
//     for head < tail {
//         if str[head] != str[tail] {
//             result = false
//             break
//         }
//         head++
//         tail--
//     }
//     return result
// }

// // 优化空间复杂度
// // defeat 88.13% users in time and 90.29% users in memory
// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	a := x % 10
// 	if a == 0 { // 这里用了trick就是非0整数的第一位肯定不是0，如果个位是0，就可以直接判断了。主要是这里降低了平均时间和空间复杂度。
// 		return x == 0
// 	}
// 	x /= 10
// 	for x > a { // 这个地方相当与只把了一半的x挪到a里面，其实已经够判断了，所以进一步降低了复杂度
// 		a = a*10 + x%10
// 		x /= 10
// 	}
// 	return a == x || a/10 == x //、
// }

// // 仿写空间复杂度优化方案
// // defeat 73% users in time and 90.29% users in memory
// func isPalindrome(x int) bool {
// 	var result bool = true
// 	if x == 0 {
// 		result = true
// 	} else if x < 0 || x%10 == 0 {
// 		result = false
// 	} else {
// 		var newnum int = 0
// 		for x > newnum {
// 			newnum = 10*newnum + x%10
// 			x = x / 10
// 		}
// 		result = (x == newnum) || (x == newnum/10)
// 	}
// 	return result
// }
