package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func New() *LinkedList {
	return &LinkedList{}
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func (l *LinkedList) Append(data int) {
	node := NewNode(data)
	l.Length++

	if l.Head == nil {
		l.Head = node
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (l *LinkedList) Prepend(data int) {
	node := NewNode(data)
	node.Next = l.Head
	l.Head = node
	l.Length++
}

func (l *LinkedList) Delete(data int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
		l.Length--
		return true
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			l.Length--
			return true
		}
		current = current.Next
	}

	return false
}

func (l *LinkedList) Search(data int) *Node {
	current := l.Head
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) String() string {
	if l.Head == nil {
		return "nil"
	}

	var parts []string
	current := l.Head
	for current != nil {
		parts = append(parts, fmt.Sprintf("%d", current.Data))
		current = current.Next
	}
	return strings.Join(parts, " -> ") + " -> nil"
}
