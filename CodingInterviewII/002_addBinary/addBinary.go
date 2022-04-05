package addBinary

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	resultLen := max(aLen, bLen) + 1
	result := make([]byte, resultLen)
	aPtr, bPtr, resultPtr := aLen-1, bLen-1, resultLen-1
	var addOne bool
	for aPtr >= 0 && bPtr >= 0 {
		if ((a[aPtr] == b[bPtr]) && addOne) || ((a[aPtr] != b[bPtr]) && !addOne) {
			result[resultPtr] = '1'
		} else {
			result[resultPtr] = '0'
		}
		addOne = (a[aPtr] == '1' || b[bPtr] == '1') && addOne || (a[aPtr] == '1' && b[bPtr] == '1')
		aPtr--
		bPtr--
		resultPtr--
	}
	for aPtr >= 0 {
		if addOne {
			if a[aPtr] == '1' {
				result[resultPtr] = '0'
				addOne = true
			} else {
				result[resultPtr] = '1'
				addOne = false
			}
		} else {
			result[resultPtr] = a[aPtr]
		}
		aPtr--
		resultPtr--
	}
	for bPtr >= 0 {
		if addOne {
			if b[bPtr] == '1' {
				result[resultPtr] = '0'
				addOne = true
			} else {
				result[resultPtr] = '1'
				addOne = false
			}
		} else {
			result[resultPtr] = b[bPtr]
		}
		bPtr--
		resultPtr--
	}
	if addOne {
		result[0] = '1'
		return string(result)
	}
	return string(result[1:])
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
