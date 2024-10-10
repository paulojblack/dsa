package ds

type LNode struct {
	Next *LNode
	Prev *LNode
	Val  int
	Key  int
}

type LRUCache struct {
	Cache    map[int]*LNode
	Head     *LNode
	Tail     *LNode
	Capacity int
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*LNode),
	}
}

func (l *LRUCache) Get(key int) int {
	node := l.Cache[key]

	if node == nil {
		return -1
	}

	l.Put(key, node.Val)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	node := l.Cache[key]

	if node == nil {
		node = createNode(key, value)
		if len(l.Cache) >= l.Capacity {
			l.detach(l.Tail)
		}
		l.prepend(node)

	} else {
		l.detach(node)
		l.prepend(node)
		node.Val = value
	}

	l.Cache[key] = node

}

func (l *LRUCache) detach(node *LNode) {
	if len(l.Cache) == 1 {
		l.Head = nil
		l.Tail = nil
	} else if node.Next == nil {
		l.Tail = node.Prev
		l.Tail.Next = nil
	} else if node.Prev == nil {
		l.Head = node.Next
		l.Head.Prev = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = nil

	delete(l.Cache, node.Key)
}

func (l *LRUCache) prepend(node *LNode) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Head.Prev = node
		node.Next = l.Head
		l.Head = node
	}
}

func createNode(key int, val int) *LNode {
	return &LNode{
		Val: val,
		Key: key,
	}
}
