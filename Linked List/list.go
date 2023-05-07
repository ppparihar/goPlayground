package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	if l.head == nil {
		l.head = &Node{value, nil}
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value, nil}
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

func (l *List) get(value int) *Node {
	if l.head == nil {
		return nil
	}
	node := l.head
	for node != nil {
		if node.value == value {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *List) print() {

	node := l.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}
