package leetcode150

// code
type LRUCache struct {
	cap  int
	size int
	data map[int]*LRUNode
	head *LRUNode
	tail *LRUNode
}

type LRUNode struct {
	key  int
	val  int
	prev *LRUNode
	next *LRUNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:  capacity,
		data: make(map[int]*LRUNode),
		head: &LRUNode{},
		tail: &LRUNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	tmp := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = tmp
	tmp.prev = node

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.val = value

		node.prev.next = node.next
		node.next.prev = node.prev

		tmp := this.head.next
		this.head.next = node
		node.prev = this.head
		node.next = tmp
		tmp.prev = node
	} else {
		newNode := &LRUNode{key: key, val: value}
		this.data[key] = newNode

		tmp := this.head.next
		this.head.next = newNode
		newNode.prev = this.head
		newNode.next = tmp
		tmp.prev = newNode

		if this.size == this.cap {
			delete(this.data, this.tail.prev.key)
			this.tail.prev = this.tail.prev.prev
			this.tail.prev.next = this.tail
		} else {
			this.size++
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
