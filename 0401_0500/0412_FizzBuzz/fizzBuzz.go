package fizzBuzz

import "strconv"

// func fizzBuzz(n int) []string {
// 	res := []string{}
// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 {
// 			if i%5 == 0 {
// 				res = append(res, "FizzBuzz")
// 			} else {
// 				res = append(res, "Fizz")
// 			}
// 		} else {
// 			if i%5 == 0 {
// 				res = append(res, "Buzz")
// 			} else {
// 				res = append(res, strconv.Itoa(i))
// 			}
// 		}
// 	}
// 	return res
// }

// append因为会涉及到自动内存分配，所以会占用更多的内存空间
// 这里因为数组空间的大小已经很清晰了，所以采用手动内存分配会审更多空间
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				res[i-1] = "FizzBuzz"
			} else {
				res[i-1] = "Fizz"
			}
		} else {
			if i%5 == 0 {
				res[i-1] = "Buzz"
			} else {
				res[i-1] = strconv.Itoa(i)
			}
		}
	}
	return res
}
