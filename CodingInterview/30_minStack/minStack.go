package minStack

type MinStack struct {
	stack    []int
	minValue []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if len(ms.minValue) == 0 || x <= ms.minValue[len(ms.minValue)-1] {
		ms.minValue = append(ms.minValue, x)
	}
}

func (ms *MinStack) Pop() {
	if ms.stack[len(ms.stack)-1] == ms.minValue[len(ms.minValue)-1] {
		ms.minValue = ms.minValue[:len(ms.minValue)-1]
	}
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Min() int {
	return ms.minValue[len(ms.minValue)-1]
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
