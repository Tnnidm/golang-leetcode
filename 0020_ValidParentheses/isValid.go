package isValid

// // my method
// // 执行用时：0 ms, 在所有 Go 提交中击败了 100.00 % 的用户
// // 内存消耗：2.1 MB, 在所有 Go 提交中击败了 15.72 % 的用户
// type byteStack struct {
// 	arr       []interface{}
// 	stackSize int
// }

// //判断是否空
// func (bs *byteStack) isEmpty() bool {
// 	return bs.stackSize == 0
// }

// //push
// func (bs *byteStack) push(element interface{}) {
// 	bs.arr = append(bs.arr, element)
// 	bs.stackSize += 1
// }

// //pop
// func (bs *byteStack) pop() interface{} {
// 	if bs.stackSize > 0 {
// 		bs.stackSize--
// 		element := bs.arr[bs.stackSize]
// 		bs.arr = bs.arr[:bs.stackSize]
// 		return element
// 	} else {
// 		return nil
// 	}
// }

// func isValid(s string) bool {
// 	result := true
// 	bs := new(byteStack)
// 	for i := range s {
// 		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
// 			bs.push(s[i])
// 		} else {
// 			if bs.isEmpty() {
// 				result = false
// 				break
// 			} else {
// 				if s[i] == ')' {
// 					if bs.pop().(byte) != '(' {
// 						result = false
// 						break
// 					}
// 				} else if s[i] == ']' {
// 					if bs.pop().(byte) != '[' {
// 						result = false
// 						break
// 					}
// 				} else if s[i] == '}' {
// 					if bs.pop().(byte) != '{' {
// 						result = false
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}
// 	if !bs.isEmpty() {
// 		result = false
// 	}
// 	return result
// }

// 如何优化内存呢？
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00 % 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 98.14 % 的用户
func isValid(s string) bool {
	var result bool = true
	strLen := len(s)
	if strLen%2 == 1 {
		return false
	} else {
		m := map[byte]byte{
			'(': ')',
			'[': ']',
			'{': '}',
		}
		stack := make([]byte, 0, strLen)
		for i := range s {
			if s[i] == '(' || s[i] == '[' || s[i] == '{' {
				stack = append(stack, s[i])
			} else {
				stackLen := len(stack)
				if stackLen == 0 {
					return false
				} else if m[stack[stackLen-1]] != s[i] {
					return false
				} else {
					stack = stack[:stackLen-1]
				}
			}
		}
		if len(stack) != 0 {
			result = false
		}
	}
	return result
}
