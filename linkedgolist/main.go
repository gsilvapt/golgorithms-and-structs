package linkedgolist

type node struct {
	data interface{} // Works for any type but requires factories for them all
	next *node
	prev *node
}

type LinkedList struct {
	head *node
	tail *node
}

func newNode(i int) *node {
	return &node{data: i}
}

func newLinkedList(i int) *LinkedList {
	var node *node = newNode(i)
	return &LinkedList{head: node, tail: node}
}

func (l *LinkedList) AddBottom(d int) *LinkedList {
	var node *node = newNode(d)
	l.tail.next = node
	return l
}

func (l *LinkedList) AddFront(d int) *LinkedList {
	var node *node = newNode(d)
	node.next = l.head
	l.head = node
	return l
}

func (l *LinkedList) Rm(d int) *LinkedList {
	var currNode *node = l.head
	for currNode != nil {
		if currNode.data == d {
			currNode.prev.next = currNode.next
			currNode = nil
		} else if currNode.next != nil {
			currNode = currNode.next
		}
	}

	return l
}

func main() {
}
