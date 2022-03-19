package maxQueue

type MaxQueue struct {
	queue    []int
	maxQueue []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (mq *MaxQueue) Max_value() int {
	if len(mq.maxQueue) == 0 {
		return -1
	}
	return mq.maxQueue[0]
}

func (mq *MaxQueue) Push_back(value int) {
	mq.queue = append(mq.queue, value)
	for len(mq.maxQueue) > 0 && mq.maxQueue[len(mq.maxQueue)-1] < value {
		mq.maxQueue = mq.maxQueue[0 : len(mq.maxQueue)-1]
	}
	mq.maxQueue = append(mq.maxQueue, value)
}

func (mq *MaxQueue) Pop_front() int {
	if len(mq.queue) < 1 {
		return -1
	}
	val := mq.queue[0]
	mq.queue = mq.queue[1:]
	if len(mq.maxQueue) > 0 && val == mq.maxQueue[0] {
		mq.maxQueue = mq.maxQueue[1:]
	}
	return val
}
