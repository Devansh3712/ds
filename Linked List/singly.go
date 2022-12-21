package main

import "fmt"

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{Data: data, Next: nil}
}

type SinglyLinkedList[T comparable] struct {
	Root *Node[T]
}

func (s *SinglyLinkedList[T]) insert(data T) {
	node := NewNode(data)
	if s.Root == nil {
		s.Root = node
		return
	}
	temp := s.Root
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
}

func (s *SinglyLinkedList[T]) remove(data T) {
	if s.Root == nil {
		return
	}
	if s.Root.Data == data {
		s.Root = s.Root.Next
	}
	temp := s.Root
	for {
		if temp.Next.Data == data {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

func (s *SinglyLinkedList[T]) search(data T) bool {
	if s.Root == nil {
		return false
	}
	temp := s.Root
	for {
		if temp == nil {
			break
		}
		if temp.Data == data {
			return true
		}
		temp = temp.Next
	}
	return false
}

func (s *SinglyLinkedList[T]) reverse() {
	if s.Root == nil || s.Root.Next == nil {
		return
	}
	var prev, curr, next *Node[T]
	prev = nil
	curr = s.Root
	next = nil
	for {
		if curr == nil {
			break
		}
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	s.Root = prev
}

func (s *SinglyLinkedList[T]) print() {
	if s.Root == nil {
		return
	}
	temp := s.Root
	for {
		if temp == nil {
			break
		}
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
}

func main() {
	var l SinglyLinkedList[int]
	for i := 0; i < 10; i++ {
		l.insert(i)
	}
	l.reverse()
	l.print()
	fmt.Println("")
}
