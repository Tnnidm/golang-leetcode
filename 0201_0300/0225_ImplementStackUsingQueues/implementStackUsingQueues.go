package implementStackUsingQueues

type MyStack struct {
	primaryQueue   []int
	assistantQueue []int
}

func Constructor() MyStack {
	myStack := MyStack{}
	return myStack
}

func (this *MyStack) Push(x int) {
	this.assistantQueue = append(this.assistantQueue, x)
	primaryQueueLen := len(this.primaryQueue)
	for i := 0; i < primaryQueueLen; i++ {
		this.assistantQueue = append(this.assistantQueue, this.primaryQueue[i])
	}
	this.primaryQueue = this.assistantQueue
	this.assistantQueue = []int{}
}

func (this *MyStack) Pop() int {
	popElement := this.primaryQueue[0]
	this.primaryQueue = this.primaryQueue[1:]
	return popElement
}

func (this *MyStack) Top() int {
	return this.primaryQueue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.primaryQueue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
