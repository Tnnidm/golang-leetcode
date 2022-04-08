package singleNumber

func singleNumber(nums []int) int {
	result := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		result |= total % 3 << i
	}
	return int(result)
}
