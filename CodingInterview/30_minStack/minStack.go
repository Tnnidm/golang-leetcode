package minStack

type MinStack struct {
	stack    []int
	minValue []int
	length   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if ms.length == 0 || x < ms.minValue[ms.length-1] {
		ms.minValue = append(ms.minValue, x)
	} else {
		ms.minValue = append(ms.minValue, ms.minValue[ms.length-1])
	}
	ms.length++
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:ms.length-1]
	ms.minValue = ms.minValue[:ms.length-1]
	ms.length--
}

func (ms *MinStack) Top() int {
	return ms.stack[ms.length-1]
}

func (ms *MinStack) Min() int {
	return ms.minValue[ms.length-1]
}

func minStack() (int, int, int) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	result1 := obj.Min()
	obj.Pop()
	result2 := obj.Top()
	result3 := obj.Min()
	return result1, result2, result3
}
