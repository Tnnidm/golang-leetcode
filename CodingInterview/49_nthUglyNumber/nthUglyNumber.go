package nthUglyNumber

// // 方法1: 使用堆和哈希表来控制插入丑数
// type minHeap []int

// func (m minHeap) Len() int {
// 	return len(m)
// }

// func (m minHeap) Less(i, j int) bool {
// 	return m[i] < m[j]
// }

// func (m minHeap) Swap(i, j int) {
// 	m[i], m[j] = m[j], m[i]
// }

// func (m *minHeap) Push(v interface{}) {
// 	*m = append(*m, v.(int))
// }

// func (m *minHeap) Pop() interface{} {
// 	mLen := len(*m)
// 	temp := (*m)[mLen-1]
// 	(*m) = (*m)[:mLen-1]
// 	return temp
// }

// func nthUglyNumber(n int) int {
// 	uglyNumber := &minHeap{}
// 	uglyNumberExit := map[int]bool{}
// 	uglyNumberExit[1] = true
// 	heap.Push(uglyNumber, 1)
// 	var thisNum, thisNumx2, thisNumx3, thisNumx5 int
// 	for n != 0 {
// 		thisNum = heap.Pop(uglyNumber).(int)
// 		thisNumx2 = 2 * thisNum
// 		thisNumx3 = 3 * thisNum
// 		thisNumx5 = 5 * thisNum
// 		if !uglyNumberExit[thisNumx2] {
// 			uglyNumberExit[thisNumx2] = true
// 			heap.Push(uglyNumber, thisNumx2)
// 		}
// 		if !uglyNumberExit[thisNumx3] {
// 			uglyNumberExit[thisNumx3] = true
// 			heap.Push(uglyNumber, thisNumx3)
// 		}
// 		if !uglyNumberExit[thisNumx5] {
// 			uglyNumberExit[thisNumx5] = true
// 			heap.Push(uglyNumber, thisNumx5)
// 		}
// 		n--
// 	}
// 	return thisNum
// }

// 方法2: 直接依次搜索下一个最小丑数来确定
func nthUglyNumber(n int) int {
	// uglyNum := make([]int, n)
	uglyNum := [1690]int{1}
	x2Index, x3Index, x5Index := 0, 0, 0
	for i := 1; i < n; i++ {
		uglyNum[i] = min(2*uglyNum[x2Index], 3*uglyNum[x3Index], 5*uglyNum[x5Index])
		if uglyNum[i] == 2*uglyNum[x2Index] {
			x2Index++
		}
		if uglyNum[i] == 3*uglyNum[x3Index] {
			x3Index++
		}
		if uglyNum[i] == 5*uglyNum[x5Index] {
			x5Index++
		}
	}
	return uglyNum[n-1]
}

func minCore(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min(x, y, z int) int {
	return minCore(minCore(x, y), z)
}
