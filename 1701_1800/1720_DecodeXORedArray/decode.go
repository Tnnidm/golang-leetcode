package decode

func decode(encoded []int, first int) []int {
	resultLen := len(encoded) + 1
	result := make([]int, resultLen)
	result[0] = first
	for i := 1; i < resultLen; i++ {
		result[i] = result[i-1] ^ encoded[i-1]
	}
	return result
}
