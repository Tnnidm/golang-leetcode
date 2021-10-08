package minStack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	stackLen := len(this.minStack)
	if stackLen == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		top := this.minStack[stackLen-1]
		this.minStack = append(this.minStack, min(x, top))
	}
}

func (this *MinStack) Pop() {
	popIndex := len(this.stack) - 1
	this.stack = this.stack[:popIndex]
	this.minStack = this.minStack[:popIndex]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
