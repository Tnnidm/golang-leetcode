package singleNumber

import "fmt"

func singleNumber(nums []int) int {
	singleOne := 0
	for _, num := range nums {
		singleOne ^= num
		fmt.Println(singleOne)
	}
	return singleOne
}
