package countPrimeSetBits

import "math/bits"

// func countPrimeSetBits(left int, right int) int {
// 	counter := 0
// 	for num := left; num <= right; num++ {
// 		oneCounter := 0
// 		for n := num; n != 0; n = n / 2 {
// 			oneCounter += n % 2
// 		}
// 		fmt.Println(num, oneCounter)
// 		primeFlag := oneCounter > 1
// 		for i := 2; i <= int(math.Sqrt(float64(oneCounter))); i++ {
// 			if oneCounter%i == 0 {
// 				primeFlag = false
// 				break
// 			}
// 		}
// 		if primeFlag {
// 			counter++
// 		}
// 	}
// 	return counter
// }

func countPrimeSetBits(left int, right int) int {
	ans := 0
	for i := left; i <= right; i++ {
		switch bits.OnesCount(uint(i)) {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			ans++
		}
	}
	return ans
}
