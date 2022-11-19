package main

import "fmt"

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) append(data any) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

func (l *LinkedList) insert(data any) {
	newNode := &Node{data: data}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) remove(data any) {
	currentNode := l.head
	if currentNode != nil && currentNode.data == data {
		l.head = currentNode.next
		currentNode = nil
		return
	}
	var previousNode *Node = nil
	for currentNode != nil && currentNode.data != data {
		previousNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode == nil {
		return
	}
	previousNode.next = currentNode.next
	currentNode = nil
}

func main() {
	l := &LinkedList{}
	l.append(1)
	l.append(2)
	l.append(3)
	l.append(4)
	l.insert(0)
	l.print()
	l.remove(5)
	fmt.Println("###############")
	l.print()
}
