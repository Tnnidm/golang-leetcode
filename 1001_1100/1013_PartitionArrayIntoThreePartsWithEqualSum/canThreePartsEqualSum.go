package canThreePartsEqualSum

// func canThreePartsEqualSum(arr []int) bool {
// 	sum := 0
// 	arrLen := len(arr)
// 	for i := 0; i < arrLen; i++ {
// 		sum += arr[i]
// 	}
// 	if sum%3 != 0 {
// 		return false
// 	} else {
// 		partSum := 0
// 		counter := 0
// 		for i := 0; i < arrLen; i++ {
// 			partSum += arr[i]
// 			if partSum == sum/3 {
// 				partSum = 0
// 				counter++
// 				if counter == 2 && i != arrLen-1 {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}
// }

func canThreePartsEqualSum(arr []int) bool {
	n := len(arr)

	//计算前缀和
	preSum := make([]int, n)
	sum := 0
	for i, val := range arr {
		preSum[i] = sum + val
		sum += val
	}

	for i := 0; i < n-2; i++ {
		if preSum[i]*3 != sum { //是否为总和的 1/3
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if (preSum[j]-preSum[i])*3 == sum { //第二部分是否为综合的 1/3
				return true
			}
		}
	}

	return false
}
