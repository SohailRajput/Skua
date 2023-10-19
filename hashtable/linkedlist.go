package hashtable

import "reflect"

/** >> LinkList **/
type Node struct {
	key   []byte
	value []byte
	next  *Node
}

type LinkedList struct {
	head *Node
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{}
}
func (l *LinkedList) insert(k []byte, v []byte) {
	newNode := &Node{
		key:   k,
		value: v,
	}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) update(k []byte, v []byte) {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(k, current.key) {
			current.value = v
			return
		}
		current = current.next
	}
}

func (l *LinkedList) search(key []byte) ([]byte, bool) {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(key, current.key) {
			return current.value, false
		}
		current = current.next
	}
	return nil, true
}
func (l *LinkedList) delete(key []byte) {
	if l.head == nil {
		return
	}
	if reflect.DeepEqual(l.head.key, key) {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if reflect.DeepEqual(current.next.key, key) {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

/** << END LinkList **/
