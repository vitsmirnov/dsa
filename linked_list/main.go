package main

type LinkedListNode[T any] struct {
	data       T
	prev, next *LinkedListNode[T]
}

type LinkedList[T any] struct {
	front, back *LinkedListNode[T]
	length      int
}

func (l *LinkedList[T]) Len() int { return l.length }

func (l *LinkedList[T]) pushBack(node *LinkedListNode[T]) *LinkedListNode[T] {
	if l.back != nil {
		l.back.next = node
	} else {
		l.front = node
	}
	l.back = node
	l.length++
	return l.back
}

func (l *LinkedList[T]) PushBack(value T) *LinkedListNode[T] {
	return l.pushBack(&LinkedListNode[T]{data: value, prev: l.back, next: nil})
}

func (l *LinkedList[T]) PopFront() T { return l.Remove(l.front) }

func (l *LinkedList[T]) Remove(node *LinkedListNode[T]) T {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.front = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.back = node.prev
	}
	l.length--
	return node.data
}

func (l *LinkedList[T]) MoveToBack(node *LinkedListNode[T]) {
	l.Remove(node)
	node.prev = l.back
	node.next = nil
	l.pushBack(node)
}
