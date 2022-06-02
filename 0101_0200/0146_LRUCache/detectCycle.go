package detectCycle

type DoubleListNode struct {
	Key   int
	Value int
	Prev  *DoubleListNode
	Next  *DoubleListNode
}

type LRUCache struct {
	key2ptr   map[int](*DoubleListNode)
	dummyHead *DoubleListNode
	dummyTail *DoubleListNode
	length    int
	capacity  int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		key2ptr:   map[int](*DoubleListNode){},
		dummyHead: &DoubleListNode{},
		dummyTail: &DoubleListNode{},
		length:    0,
		capacity:  capacity}
	lru.dummyHead.Next = lru.dummyTail
	lru.dummyTail.Prev = lru.dummyHead
	return lru
}

func (this *LRUCache) Get(key int) int {
	if ptr, ok := this.key2ptr[key]; ok {
		deleteNode(ptr)
		addNodeBehind(this.dummyHead, ptr)
		return ptr.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if ptr, ok := this.key2ptr[key]; ok {
		ptr.Value = value
		deleteNode(ptr)
		addNodeBehind(this.dummyHead, ptr)
	} else {
		if this.length >= this.capacity {
			delete(this.key2ptr, this.dummyTail.Prev.Key)
			deleteNode(this.dummyTail.Prev)
			this.length--
		}
		newNode := &DoubleListNode{Key: key, Value: value}
		addNodeBehind(this.dummyHead, newNode)
		this.key2ptr[key] = newNode
		this.length++
	}
}

func deleteNode(node *DoubleListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func addNodeBehind(prev, node *DoubleListNode) {
	node.Next = prev.Next
	prev.Next.Prev = node
	prev.Next = node
	node.Prev = prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
