package main

/*type Node struct {
	data any
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (dl *DoublyLinkedList) append(data any) {
	newNode := &Node{data: data}
	if dl.head == nil {
		dl.head = newNode
		return
	}
	currentNode := dl.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
	newNode.prev = currentNode
}

func (dl *DoublyLinkedList) insert(data any) {
	newNode := &Node{data: data}
	if dl.head == nil {
		dl.head = newNode
		return
	}

	dl.head.prev = newNode
	newNode.next = dl.head
	dl.head = newNode
}

func (dl *DoublyLinkedList) print() {
	currentNode := dl.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (dl *DoublyLinkedList) remove(data any) {
	currentNode := dl.head
	if currentNode != nil && currentNode.data == data {
		if currentNode.next == nil {
			currentNode = nil
			dl.head = nil
			return
		} else {
			nextNode := currentNode.next
			nextNode.prev = nil
			currentNode = nil
			dl.head = nextNode
			return
		}
	}

	for currentNode != nil && currentNode.data != data {
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return
	}
	if currentNode.next == nil {
		prev := currentNode.prev
		prev.next = nil
		currentNode = nil
		return
	} else {
		nextNode := currentNode.next
		prevNode := currentNode.prev
		prevNode.next = nextNode
		nextNode.prev = prevNode
		currentNode = nil
		return
	}
}

func (dl *DoublyLinkedList) reverseIterative() {
	var previousNode *Node = nil
	currentNode := dl.head
	for currentNode != nil {
		previousNode = currentNode.prev
		currentNode.prev = currentNode.next
		currentNode.next = previousNode

		currentNode = currentNode.prev
	}
	if previousNode != nil {
		dl.head = previousNode.prev
	}
}

func main() {
	dl := &DoublyLinkedList{}
	dl.append(0)
	dl.append(1)
	dl.append(2)
	dl.append(3)
	dl.print()
	fmt.Println("######## Reverse iter")
	dl.reverseIterative()
	dl.print()
}
*/
