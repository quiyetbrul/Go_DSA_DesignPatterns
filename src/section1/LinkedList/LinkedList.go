package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Push(value int) {
	n := &Node{
		value: value,
	}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.next = n
		n.prev = l.Tail
	}

	l.Tail = n
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	fmt.Println(l.First().value)

	n := l.First()
	for n != nil {
		fmt.Println(n.value)
		n = n.Next()
	}

	n = l.Last()
	for n != nil {
		fmt.Println(n.value)
		n = n.Prev()
	}
}
