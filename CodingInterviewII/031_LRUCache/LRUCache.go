package LRUCache

type BiDiListNode struct {
	key   int
	value int
	prev  *BiDiListNode
	next  *BiDiListNode
}

type LRUCache struct {
	ptrMap    map[int](*BiDiListNode)
	dummyHead *BiDiListNode
	dummyTail *BiDiListNode
	size      int
	capacity  int
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{map[int](*BiDiListNode){}, &BiDiListNode{}, &BiDiListNode{}, 0, capacity}
	lc.dummyHead.next = lc.dummyTail
	lc.dummyTail.prev = lc.dummyHead
	return lc
}

func (lc *LRUCache) Get(key int) int {
	if ptr, ok := lc.ptrMap[key]; ok {
		lc.moveToHead(ptr)
		return ptr.value
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if ptr, ok := lc.ptrMap[key]; ok {
		ptr.value = value
		lc.moveToHead(ptr)
	} else {
		newNode := &BiDiListNode{key, value, nil, nil}
		lc.ptrMap[key] = newNode
		lc.addToHead(newNode)
		if lc.size < lc.capacity {
			lc.size++
		} else {
			odd := lc.removeTail()
			delete(lc.ptrMap, odd.key)
		}
	}
}

func (lc *LRUCache) addToHead(node *BiDiListNode) {
	node.prev = lc.dummyHead
	node.next = lc.dummyHead.next
	lc.dummyHead.next.prev = node
	lc.dummyHead.next = node
}

func (lc *LRUCache) removeNode(node *BiDiListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lc *LRUCache) moveToHead(node *BiDiListNode) {
	lc.removeNode(node)
	lc.addToHead(node)
}

func (lc *LRUCache) removeTail() *BiDiListNode {
	node := lc.dummyTail.prev
	lc.removeNode(node)
	return node
}
