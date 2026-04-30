package main

type LinkedListNode[T any] struct {
	Data       T
	Prev, Next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	front, back *LinkedListNode[T]
	length      int
}

func (l *LinkedList[T]) Len() int                  { return l.length }
func (l *LinkedList[T]) Front() *LinkedListNode[T] { return l.front }
func (l *LinkedList[T]) Back() *LinkedListNode[T]  { return l.back }

func (l *LinkedList[T]) pushBack(node *LinkedListNode[T]) *LinkedListNode[T] {
	if l.back != nil {
		l.back.Next = node
	} else {
		l.front = node
	}
	l.back = node
	l.length++
	return l.back
}

func (l *LinkedList[T]) PushBack(value T) *LinkedListNode[T] {
	return l.pushBack(&LinkedListNode[T]{Data: value, Prev: l.back, Next: nil})
}

func (l *LinkedList[T]) PopFront() T { return l.Remove(l.front) }

func (l *LinkedList[T]) Remove(node *LinkedListNode[T]) T {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.front = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.back = node.Prev
	}
	l.length--
	return node.Data
}

func (l *LinkedList[T]) MoveToBack(node *LinkedListNode[T]) {
	l.Remove(node)
	node.Prev = l.back
	node.Next = nil
	l.pushBack(node)
}

// todo: iterator
