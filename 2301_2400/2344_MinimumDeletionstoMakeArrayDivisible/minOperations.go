package minOperations

import "sort"

func minOperations(nums []int, numsDivide []int) int {
	gcd := numsDivide[0]
	for i := 1; i < len(numsDivide); i++ {
		gcd = findGCD(gcd, numsDivide[i])
	}
	sort.Ints(nums)
	for i := range nums {
		if gcd%nums[i] == 0 {
			return i
		}
	}
	return -1
}

func findGCD(x int, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
