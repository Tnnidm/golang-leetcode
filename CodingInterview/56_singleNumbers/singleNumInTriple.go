package singleNumbers

func singleNumberInTriple(nums []int) int {
	assistArr := make([]int, 32)
	for _, num := range nums {
		index := 0
		for num != 0 {
			assistArr[index] += num & 1
			num >>= 1
			index++
		}
	}
	result := 0
	for k, index := 1, 0; index < 32; k, index = k<<1, index+1 {
		result += k * (assistArr[index] % 3)
	}
	return result
}
