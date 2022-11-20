package main

import "fmt"

type node struct {
	Value int
}

func (n *node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func main() {
	s := NewStack()
	s.Push(&node{3})
	s.Push(&node{5})
	s.Push(&node{7})
	s.Push(&node{9})
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}
