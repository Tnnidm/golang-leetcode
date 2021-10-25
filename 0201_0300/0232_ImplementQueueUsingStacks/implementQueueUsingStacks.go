package implementQueueUsingStacks

type MyQueue struct {
	primaryStack   []int
	assistantStack []int
}

func Constructor() MyQueue {
	q := MyQueue{}
	return q
}

func (this *MyQueue) Push(x int) {
	this.assistantStack = []int{}
	for i := len(this.primaryStack) - 1; i >= 0; i-- {
		this.assistantStack = append(this.assistantStack, this.primaryStack[i])
	}
	this.assistantStack = append(this.assistantStack, x)
	this.primaryStack = []int{}
	for i := len(this.assistantStack) - 1; i >= 0; i-- {
		this.primaryStack = append(this.primaryStack, this.assistantStack[i])
	}
}

func (this *MyQueue) Pop() int {
	primaryStackLen := len(this.primaryStack)
	element := this.primaryStack[primaryStackLen-1]
	this.primaryStack = this.primaryStack[:primaryStackLen-1]
	return element
}

func (this *MyQueue) Peek() int {
	primaryStackLen := len(this.primaryStack)
	element := this.primaryStack[primaryStackLen-1]
	return element
}

func (this *MyQueue) Empty() bool {
	return len(this.primaryStack) == 0
}
