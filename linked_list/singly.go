package main

import (
	"cmp"
	"fmt"
)

// todo: review, iterator

type LinkedListNode1[T any] struct {
	Data T
	Next *LinkedListNode1[T]
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

func (this *LinkedList1[T]) Len() int                  { return this.length }
func (this *LinkedList1[T]) Head() *LinkedListNode1[T] { return this.head }
func (this *LinkedList1[T]) Tail() *LinkedListNode1[T] { return this.tail }

func (this *LinkedList1[T]) Get(index int) T {
	if !(0 <= index && index < this.length) {
		panic(fmt.Sprintf("LinkedList1: index (%v) out of range [0-%v)", index, this.length))
	}
	if index == 0 {
		return this.head.Data
	} else if index == this.length-1 {
		return this.tail.Data
	} else {
		return this.getNode(index).Data
	}
}

func (this *LinkedList1[T]) AddAtHead(data T) {
	this.head = &LinkedListNode1[T]{Data: data, Next: this.head}
	if this.tail == nil {
		this.tail = this.head
	}
	this.length++
}

func (this *LinkedList1[T]) AddAtTail(data T) {
	if this.tail == nil {
		this.tail = &LinkedListNode1[T]{Data: data, Next: nil}
		this.head = this.tail
	} else {
		this.tail.Next = &LinkedListNode1[T]{Data: data, Next: nil}
		this.tail = this.tail.Next
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
		prev.Next = &LinkedListNode1[T]{Data: data, Next: prev.Next}
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
		this.head = this.head.Next
	} else {
		prev := this.getNode(index - 1)
		prev.Next = prev.Next.Next
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
		cur = cur.Next
	}
	return cur
}

// todo: tests for sorting
// change to interface?

func SortList[T cmp.Ordered](head *LinkedListNode1[T]) *LinkedListNode1[T] {
	// time: O(n log n), space: O(log n)

	if head == nil || head.Next == nil {
		return head
	}

	mid := CutListInHalf(head)
	head = SortList(head)
	mid = SortList(mid)
	return mergeSortedLists(head, mid)
}

func mergeSortedLists[T cmp.Ordered](list1, list2 *LinkedListNode1[T]) *LinkedListNode1[T] {
	// time: O(n), space: O(1)

	preHead := &LinkedListNode1[T]{Next: nil}
	cur := preHead
	for ; list1 != nil && list2 != nil; cur = cur.Next {
		if list1.Data < list2.Data {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return preHead.Next
}

func CutListInHalf[T any](head *LinkedListNode1[T]) *LinkedListNode1[T] {
	// time: O(n), space: O(1)

	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	predSlow := (*LinkedListNode1[T])(nil)
	for fast != nil && fast.Next != nil {
		predSlow = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	predSlow.Next = nil
	return slow
}
