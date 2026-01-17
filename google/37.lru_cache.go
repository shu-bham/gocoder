package google

type DLNode struct {
	Key, Val   int
	Prev, Next *DLNode
}
type LRUCache struct {
	Head     *DLNode
	Tail     *DLNode
	KeyMap   map[int]*DLNode
	Capacity int
}

// renamed to use other Constructor
func LRUCacheConstructor(capacity int) LRUCache {
	head := &DLNode{}
	tail := &DLNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Head:     head,
		Tail:     tail,
		KeyMap:   make(map[int]*DLNode),
		Capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.KeyMap[key]
	if ok {
		deleted := l.deleteNode(node)
		l.insertNode(deleted)
		return node.Val
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.KeyMap[key]
	if ok {
		deleted := l.deleteNode(node)
		deleted.Val = value
		l.insertNode(deleted)
	} else {
		if len(l.KeyMap) == l.Capacity {
			l.deleteNode(l.Head.Next)
			l.insertNode(newNode(key, value))
		} else {
			l.insertNode(newNode(key, value))
		}
	}
}

func newNode(key, value int) *DLNode {
	return &DLNode{
		Key:  key,
		Val:  value,
		Prev: nil,
		Next: nil,
	}
}

func (l *LRUCache) insertNode(node *DLNode) {
	// insert at tail
	prev := l.Tail.Prev
	node.Next = l.Tail
	l.Tail.Prev = node
	node.Prev = prev
	prev.Next = node
	l.KeyMap[node.Key] = node
}

func (l *LRUCache) deleteNode(node *DLNode) *DLNode {
	next := node.Next
	prev := node.Prev
	prev.Next = next
	next.Prev = prev
	node.Prev = nil
	node.Next = nil
	delete(l.KeyMap, node.Key)
	return node
}
