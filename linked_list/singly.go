package main

import "fmt"

// todo: review, iterator

type LinkedListNode1[T any] struct {
	data T
	next *LinkedListNode1[T]
}

type LinkedList1[T any] struct {
	head   *LinkedListNode1[T]
	tail   *LinkedListNode1[T]
	length int
}

func MakeLinkedList1[T any]() *LinkedList1[T] {
	return &LinkedList1[T]{
		head:   nil,
		tail:   nil,
		length: 0}
}

func (this *LinkedList1[T]) Len() int { return this.length }

func (this *LinkedList1[T]) Get(index int) T {
	if !(0 <= index && index < this.length) {
		panic(fmt.Sprintf("LinkedList1: index (%v) out of range [0-%v)", index, this.length))
	}
	if index == 0 {
		return this.head.data
	} else if index == this.length-1 {
		return this.tail.data
	} else {
		return this.getNode(index).data
	}
}

func (this *LinkedList1[T]) AddAtHead(data T) {
	this.head = &LinkedListNode1[T]{data: data, next: this.head}
	if this.tail == nil {
		this.tail = this.head
	}
	this.length++
}

func (this *LinkedList1[T]) AddAtTail(data T) {
	if this.tail == nil {
		this.tail = &LinkedListNode1[T]{data: data, next: nil}
		this.head = this.tail
	} else {
		this.tail.next = &LinkedListNode1[T]{data: data, next: nil}
		this.tail = this.tail.next
	}
	this.length++
}

func (this *LinkedList1[T]) AddAtIndex(index int, data T) {
	if !(0 <= index && index <= this.length) {
		return
	}
	if index == 0 {
		this.AddAtHead(data)
	} else if index == this.length {
		this.AddAtTail(data)
	} else {
		prev := this.getNode(index - 1)
		prev.next = &LinkedListNode1[T]{data: data, next: prev.next}
		this.length++
	}
}

func (this *LinkedList1[T]) DeleteAtIndex(index int) {
	if !(0 <= index && index < this.length) {
		return
	}
	if index == 0 {
		if this.tail == this.head {
			this.tail = nil
		}
		this.head = this.head.next
	} else {
		prev := this.getNode(index - 1)
		prev.next = prev.next.next
		if index == this.length-1 {
			this.tail = prev
		}
	}
	this.length--
}

func (this *LinkedList1[T]) getNode(index int) *LinkedListNode1[T] {
	if !(0 <= index && index < this.length) {
		return nil
	}
	cur := this.head
	for range index {
		cur = cur.next
	}
	return cur
}
