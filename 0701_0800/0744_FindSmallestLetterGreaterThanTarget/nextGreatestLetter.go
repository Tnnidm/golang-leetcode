package nextGreatestLetter

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}

// func nextGreatestLetter(letters []byte, target byte) byte {
// 	l, r := 0, len(letters)-1
// 	for l <= r {
// 		mid := l + (r-l)/2
// 		if target < letters[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return letters[l%len(letters)]
// }
